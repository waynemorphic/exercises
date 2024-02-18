package main

import "math"

// If a number is prime, number i.e, if a number is positive number and has 1 and itself only 
// as the positive divisors
// i.e, -1 -> false
// 0 -> false   6 -> false
// 1 -> false   7 -> true
// 2 -> true    8 -> false
// 3 -> true	9 -> false
// 4 -> false	10 -> false
// 5 -> true

func IsPrime(num int) bool {
	if num <= 1 {
		return false
	}

	// Optimize the loop
	var j = int(math.Sqrt(float64(num))) + 1

	for i := 2; i < j; i++ {
		if num % i == 0{
			return false
		}
	}
	return true
}