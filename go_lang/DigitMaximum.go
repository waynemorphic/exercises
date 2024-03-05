package main

// Write a function that given an array A consisting of N integers, returns the maximum among all one digit integers

// Example, given array A = [-6, -91, 1011, -100, 84, -22, 0, 1, 473]

// the function should return 1

import(
	"math"
	"fmt"
)

func DigitMaximum(arr []int) int{
	var res int = 0
	
	for _, j := range(arr){
		if j > 0 && j < 10 {
			res = int(math.Max(float64(res), float64(j)))
		}
	}
	return res
}

func main(){
	fmt.Println(DigitMaximum([]int {-6, -91, 1011, -100, 84, -22, 0, 1, 473}))
	fmt.Println(DigitMaximum([]int {1,2,3}))
	fmt.Println(DigitMaximum([]int {-1,-2,-3}))
}

// go run DigitMaximum.go