package main

import (
	"fmt"
)
// If a number is prime, number i.e, if a number is positive number and has 1 and itself only 
// as the positive divisors
// i.e, -1 -> false
// 0 -> false   6 -> false
// 1 -> false   7 -> true
// 2 -> true    8 -> false
// 3 -> true	9 -> false
// 4 -> false	10 -> false
// 5 -> true

func isPrime(num int) bool {
	if num <= 1 {
		return false
	}

	for i := 2; i < num; i++ {
		if num % i == 0{
			return false
		}
	}
	return true
}

func main(){
	for i := -1; i <= 10; i++ {
		fmt.Println(isPrime(i))
	}
}


// Create and initialize a go module -> go mod init <the package name path>
// Run a Go function i.e, main() -> go run main.go