package main

/*
A valid password must meet the following criteria:

    - At least 5 characters long but no more than 12 characters.
    - Contains at least one uppercase letter.
    - Contains at least one digit.


Implement the isValidPassword function by looping through each character in the password string.
Make sure the password is long enough and includes at least one uppercase letter and one digit.
*/

func isValidPassword(password string) bool {
	length := len(password)
	hasUpper := false
	hasDigit := false
	hasFittingLength := false

	if length >= 5 && length <= 12 {
		hasFittingLength = true
	}

	for i := 0; i < length; i++ {
		char := password[i]

		if char >= 'A' && char <= 'Z' {
			hasUpper = true
		}

		if char >= '0' && char <= '9' {
			hasDigit = true
		}
	}
	return hasUpper && hasDigit && hasFittingLength
}

func isValidPasswordAlternate(password string) bool {
	digits := "0123456789"
	uppercase := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	foundDigit := false
	foundUpper := false

	if len(password) < 5 || len(password) > 12 {
		return false
	}

	for _, val := range password {

		for _, i := range digits {
			if val == i {
				foundDigit = true
				break
			}
		}

		for _, j := range uppercase {
			if val == j {
				foundUpper = true
				break
			}
		}
	}

	if foundDigit == foundUpper {
		return true
	}

	return false
}

func main() {}
