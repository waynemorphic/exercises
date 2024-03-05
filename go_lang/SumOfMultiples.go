package main

// Write a function that, given an array A of N integers, returns the sum of all integers which are multiples of 4

// For example, given array A as follows:

// [-6, -91, 1011, -100, 84, -22, 0, 1, 473]

// the function should return -16

import (
	"fmt"
)

func SumOfmultiples(arr []int) int{
	res := 0

	for _, j := range(arr) {
		if j % 4 == 0{
			res += j
		}
	}
	return res
}

func main(){
	fmt.Println(SumOfmultiples([]int {-6, -91, 1011, -100, 84, -22, 0, 1, 473}))
}