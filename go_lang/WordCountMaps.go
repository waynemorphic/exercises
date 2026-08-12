package main

/*
Exercise is from official Golang tour

Implement WordCount. It should return a map of the counts of each “word” in the string s.
The wc.Test function runs a test suite against the provided function and prints success or failure.

You might find strings.Fields helpful.
*/

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	resMap := make(map[string]int)
	stringFields := strings.Fields(s)

	for idx, words := range stringFields {
		count := 0
		for i := len(stringFields) - 1; i >= 0; i-- {
			if stringFields[i] == stringFields[idx] {
				count++
			}
		}
		resMap[words] = count
	}

	return resMap
}

func main() {
	// go mod tidy -> Installs the TestSuite import
	// go run WordCountMap.go
	wc.Test(WordCount)
}
