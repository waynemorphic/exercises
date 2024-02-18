package main

import (
	"fmt"
)


func main(){
	// IsPrime function call
	for i := -1; i <= 10; i++ {
		fmt.Println(IsPrime(i))
	}

	// ArrayDiff function call
	var emptyArray []int
	fmt.Println(ArrayDiff([]int {1,2}, []int {1}))
	fmt.Println(ArrayDiff(emptyArray, []int {1,2,3,4}))
	fmt.Println(ArrayDiff([]int {1, 2, 2, 2, 3}, []int {2}))
}


// Create and initialize a go module -> go mod init <the package name path>
// Run a Go function i.e, main() -> go run main.go <file to run>