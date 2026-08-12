package main

import "errors"

/*
Complete the getUserMap function.
It takes a slice of names and a slice of phone numbers, and returns a map of name -> user structs and an error.
A user struct just contains a user's name and phone number. The first element in the names slice pairs with the first
phone number, and so on.

If the length of names and phoneNumbers is not equal, return an error with the string "invalid sizes".
*/

func main() {}

func getUserMap(names []string, phoneNumbers []string) (map[string]user, error) {
	if len(names) != len(phoneNumbers) {
		return nil, errors.New("invalid sizes")
	}

	users := make(map[string]user)

	for i, name := range names {
		users[name] = user{
			name:        name,
			phoneNumber: phoneNumbers[i],
		}
	}

	return users, nil
}

type user struct {
	name        string
	phoneNumber string
}

/*
myMap := map[string]int
Example: myMap[string]int{ "John": 1, "James": 2 }

1. Insert an element in a map
- key is a string, elem is an int
myMap[key] = elem

2. Get an element from a map
elem = myMap[key]

3. Delete an element from a map
delete(myMap, key)

4. Check if a key exists in a map
elem, ok := myMap[key]

*/
