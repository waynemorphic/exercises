package main

import (
	"context"
	"errors"
	"fmt"
	"math"
	"math/rand/v2"
	"time"
)

type Retrier struct {
	attempts  int
	backoff   BackoffFunc
	fraction  float64
	predicate Predicate
	hook      Hook
	timeout   time.Duration
	sleep     func(ctx context.Context, wait time.Duration) error
}

type BackoffFunc func(attempt int) time.Duration
type Predicate func(error) bool
type Hook func(attempt int, err error, wait time.Duration)
type Option func(*Retrier) error

func New(options ...Option) (*Retrier, error) {
	retrier := &Retrier{
		attempts: 3,
		backoff: func(attempt int) time.Duration {
			return 100 * time.Millisecond
		},
		sleep: sleepWithContext,
	}

	for _, option := range options {
		if option == nil {
			return nil, errors.New("retry: nil option")
		}
		if err := option(retrier); err != nil {
			return nil, err
		}
	}
	return retrier, nil
}

/*
WithMaxAttempts
Total attempts, including the first. Must be ≥ 1, otherwise return an error.
*/
func WithMaxAttempts(attempts int) Option {
	return func(r *Retrier) error {
		if attempts < 1 {
			return errors.New("retry: attempts < 1")
		}
		r.attempts = attempts
		return nil
	}
}

/*
WithBackoff
type BackoffFunc func(attempt int) time.Duration. Decides the wait before the next attempt.
*/
func WithBackoff(backoff BackoffFunc) Option {
	return func(r *Retrier) error {
		if backoff == nil {
			return errors.New("retry: backoff == nil")
		}
		r.backoff = backoff
		return nil
	}
}

/*
WithConstantBackoff
Convenience option built on WithBackoff.
*/
func WithConstantBackoff(d time.Duration) Option {
	return func(r *Retrier) error {
		if d < time.Second {
			return errors.New("retry: backoff < time.Second")
		}
		return WithBackoff(func(attempt int) time.Duration {
			return d
		})(r)
	}
}

/*
WithExponentialBackoff
Doubles each attempt, capped at max.
*/
func WithExponentialBackoff(base, max time.Duration) Option {
	return func(r *Retrier) error {
		if base < 0 || max < 0 {
			return errors.New("retry: duration should not be < 0")
		}

		if base > max {
			return errors.New("retry: base duration should be less than max")
		}

		return WithBackoff(func(attempt int) time.Duration {
			delay := base

			for i := 0; i < attempt; i++ {
				if delay > max/2 {
					return max
				}
				delay *= 2
			}
			return delay
		})(r)
	}
}

/*
WithJitter
Randomizes the delay by ±fraction (0 to 1). Must validate the range.
*/
func WithJitter(fraction float64) Option {
	return func(r *Retrier) error {
		if fraction < 0.0 || fraction > 1.0 || math.IsNaN(fraction) {
			return errors.New("retry: invalid fraction")
		}
		r.fraction = fraction
		return nil
	}
}

/*
WithRetryIf
Only retry when the predicate returns true. Default: retry on any non-nil error.
*/
func WithRetryIf(pred func(err error) bool) Option {
	return func(r *Retrier) error {
		if pred == nil {
			return errors.New("retry: pred == nil")
		}
		r.predicate = pred
		return nil
	}
}

/*
WithOnRetry
Called before each wait, useful for logging and metrics.
*/
func WithOnRetry(hook func(attempt int, err error, wait time.Duration)) Option {
	return func(r *Retrier) error {
		if hook == nil {
			return errors.New("retry: hook == nil")
		}
		r.hook = hook
		return nil
	}
}

/*
WithTimeout
Overall deadline across all attempts.
*/
func WithTimeout(d time.Duration) Option {
	return func(r *Retrier) error {
		if d < 0 {
			return errors.New("retry: duration must not be negative")
		}
		r.timeout = d
		return nil
	}
}

func (r *Retrier) Do(ctx context.Context, fn func(ctx context.Context) error) error {
	if fn == nil {
		return errors.New("retry: function must not ne nil")
	}

	if ctx == nil {
		return errors.New("retry: nil context cannot be passed")
	}

	if r.timeout > 0 {
		timeoutCtx, cancel := context.WithTimeout(ctx, r.timeout)
		defer cancel()
		ctx = timeoutCtx
	}

	var lastErr error

	for attempt := 1; attempt <= r.attempts; attempt++ {
		// Another attempt cannot be started after cancellation
		if err := ctx.Err(); err != nil {
			return errors.Join(err, lastErr)
		}

		lastErr = fn(ctx)
		if lastErr == nil {
			return nil
		}

		// A nil predicate means retry on any failure by default
		if r.predicate != nil && !r.predicate(lastErr) {
			return lastErr
		}

		// No backoff or hook after final attempt
		if attempt == r.attempts {
			return fmt.Errorf("retry: failed after %d attempts: %w", attempt, lastErr)
		}

		if err := ctx.Err(); err != nil {
			return errors.Join(err, lastErr)
		}

		wait := r.backoff(attempt)
		if wait < 0 {
			return fmt.Errorf("retry: backoff returned a negative duration: %w", lastErr)
		}

		if r.fraction > 0 {
			// A fraction of 0.2 produces a multiplier in [0.8, 1.2).
			multiplier := 1 + (2*rand.Float64()-1)*r.fraction
			jittered := float64(wait) * multiplier

			// Saturate rather than overflow time.Duration
			const maxDuration time.Duration = math.MaxInt64

			if jittered >= float64(maxDuration) {
				wait = maxDuration
			} else {
				wait = time.Duration(jittered)
			}
		}

		if r.hook != nil {
			r.hook(attempt, lastErr, wait)
		}

		if err := r.sleep(ctx, wait); err != nil {
			return errors.Join(err, lastErr)
		}
	}
	return errors.New("retry: max attempts must be at least 1")
}

func sleepWithContext(ctx context.Context, wait time.Duration) error {
	if err := ctx.Err(); err != nil {
		return err
	}

	if wait <= 0 {
		return nil
	}

	timer := time.NewTimer(wait)
	defer timer.Stop()

	select {
	case <-ctx.Done():
		return ctx.Err()
	case <-timer.C:
		return nil
	}
}

func main() {
	calls := 0

	r, err := New(
		WithMaxAttempts(5),
		WithExponentialBackoff(200*time.Millisecond, 2*time.Second),
		WithJitter(0.2),
		WithOnRetry(func(attempt int, err error, wait time.Duration) {
			fmt.Printf("attempt %d failed: %v (retrying in %s)\n", attempt, err, wait)
		}),
	)
	if err != nil {
		fmt.Println("setup error:", err)
		return
	}

	err = r.Do(context.Background(), func(ctx context.Context) error {
		calls++
		if calls < 3 {
			return errors.New("flaky failure")
		}
		return nil
	})
	fmt.Println("result:", err, "after", calls, "calls")
}
