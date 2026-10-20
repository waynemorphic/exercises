# Challenge: Build a Configurable `Retrier` with Functional Options

## Scenario

Your team keeps copy-pasting retry loops around flaky calls: HTTP requests, S3 uploads, message publishing. You're asked
to write a small reusable `retry` package. Different call sites need different behavior, so the constructor must stay
stable while new knobs get added over time. That is the situation Go **functional options** are designed for.

## Requirements

Create a package `retry` with this core API:

```go
type Retrier struct { /* unexported fields */ }

type Option func (*Retrier) error

func New(opts ...Option) (*Retrier, error)

func (r *Retrier) Do(ctx context.Context, fn func (ctx context.Context) error) error
```

### Options to implement

| Option                                                               | Behavior                                                                                      |
|----------------------------------------------------------------------|-----------------------------------------------------------------------------------------------|
| `WithMaxAttempts(n int)`                                             | Total attempts, including the first. Must be ≥ 1, otherwise return an error.                  |
| `WithBackoff(b BackoffFunc)`                                         | `type BackoffFunc func(attempt int) time.Duration`. Decides the wait before the next attempt. |
| `WithConstantBackoff(d time.Duration)`                               | Convenience option built on `WithBackoff`.                                                    |
| `WithExponentialBackoff(base, max time.Duration)`                    | Doubles each attempt, capped at `max`.                                                        |
| `WithJitter(fraction float64)`                                       | Randomizes the delay by ±fraction (0 to 1). Must validate the range.                          |
| `WithRetryIf(pred func(error) bool)`                                 | Only retry when the predicate returns true. Default: retry on any non-nil error.              |
| `WithOnRetry(hook func(attempt int, err error, wait time.Duration))` | Called before each wait, useful for logging and metrics.                                      |
| `WithTimeout(d time.Duration)`                                       | Overall deadline across all attempts.                                                         |

### Defaults

If no options are passed, `New()` should give: 3 attempts, 100ms constant backoff, no jitter, retry on any error, no
hook, no overall timeout.

### Behavior rules for `Do`

1. Return `nil` as soon as `fn` succeeds.
2. If the error fails the `WithRetryIf` predicate, return it immediately without retrying.
3. If the context is cancelled or times out during a wait, stop and return `ctx.Err()` wrapped with the last error from
   `fn`. Use `errors.Join` or `%w`, so both are inspectable with `errors.Is`.
4. After the final attempt, return an error that wraps the last failure and says how many attempts were made.
5. Never `time.Sleep`. Use a `select` on `ctx.Done()` and a timer so cancellation is respected.

## Example usage your implementation must support

```go
r, err := retry.New(
retry.WithMaxAttempts(5),
retry.WithExponentialBackoff(200*time.Millisecond, 5*time.Second),
retry.WithJitter(0.2),
retry.WithRetryIf(func (err error) bool {
var ne net.Error
return errors.As(err, &ne) && ne.Timeout()
}),
retry.WithOnRetry(func (attempt int, err error, wait time.Duration) {
log.Printf("attempt %d failed: %v (retrying in %s)", attempt, err, wait)
}),
)
if err != nil {
log.Fatal(err)
}

err = r.Do(ctx, func(ctx context.Context) error {
return uploadToS3(ctx, file)
})
```

## Constraints

- No exported struct fields on `Retrier`. All configuration goes through options.
- Options must be order-independent where reasonable. For example, `WithJitter` before `WithBackoff` should behave the
  same as after.
- Invalid options must surface as an error from `New`, not a panic and not silently ignored.
- Each option should be a small closure, with no giant switch statement or config struct passed in.
- Make the sleep source injectable for tests. A small unexported `sleep` field, set by an unexported option or by a test
  in the same package, is enough, so your tests don't actually wait.

## What to test

- Defaults are applied when `New()` is called with no options
- `WithMaxAttempts(0)` and `WithJitter(1.5)` return errors
- `fn` that fails twice then succeeds returns `nil` after exactly 3 calls
- A non-retryable error stops after 1 call
- Context cancelled mid-backoff returns promptly
- `WithOnRetry` is called the right number of times with correct attempt numbers
- Exponential backoff produces `200ms, 400ms, 800ms, …` capped at `max` (with jitter off)

## Stretch goals

1. Add `WithClock(clock Clock)` so time can be fully faked in tests.
2. Add a generic helper
   `func DoValue[T any](ctx context.Context, r *Retrier, fn func(ctx context.Context) (T, error)) (T, error)`.
3. Add a `WithBudget(tokens int)` option that limits retries shared across many `Do` calls (a retry budget), which
   requires thinking about concurrency safety.
4. Write a benchmark showing that `New(...)` with 8 options doesn't allocate more than you expect.

## Things to think about when you're done

- Why return `error` from `Option` instead of keeping `type Option func(*Retrier)`? What do you gain, and what do you
  pay?
- If you added a field to `Retrier` next month, which existing callers would have to change?
- Where would a plain config struct have been simpler? Functional options aren't always the right tool.
