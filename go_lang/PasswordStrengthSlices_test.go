package main

import (
	"fmt"
	"testing"
)

func TestIsValidPassword(t *testing.T) {
	type testCase struct {
		password string
		isValid  bool
	}

	runCases := []testCase{
		{"Pass123", true},
		{"pas", false},
		{"Password", false},
		{"123456", false},
	}

	submitCases := append(runCases, []testCase{
		{"VeryLongPassword1", false},
		{"Short", false},
		{"1234short", false},
		{"Short5", true},
		{"P4ssword", true},
		{"AA0Z9", true},
	}...)

	testCases := runCases

	skipped := len(submitCases) - len(testCases)

	passCount := 0
	failCount := 0

	for i, test := range testCases {
		t.Run(fmt.Sprintf("TestCase%d", i+1), func(t *testing.T) {
			result := isValidPasswordAlternate(test.password)
			if result != test.isValid {
				failCount++
				t.Errorf(`---------------------------------
Password:  "%s"
Expecting: %v
Actual:    %v
Fail
`, test.password, test.isValid, result)
			} else {
				passCount++
				fmt.Printf(`---------------------------------
Password:  "%s"
Expecting: %v
Actual:    %v
Pass
`, test.password, test.isValid, result)
			}
		})
	}

	fmt.Println("---------------------------------")
	if skipped > 0 {
		fmt.Printf("%d passed, %d failed, %d skipped\n", passCount, failCount, skipped)
	} else {
		fmt.Printf("%d passed, %d failed\n", passCount, failCount)
	}
}
