package main

/*
Each time a user is sent a message, their username is logged in a slice. We want a more efficient way to count how many
messages each user received.

Implement the updateCounts function. It takes as input:

    messagedUsers: a slice of strings.
    validUsers: a map of string -> int.

It should update the validUsers map with the number of times each user has received a message. Each string in the slice
is a username, but they may not be valid. Only update the message count of valid users.

So, if "benji" is in the map and appears in the slice 3 times, the key "benji" in the map should have the value 3.
*/

func updateCounts(messagedUsers []string, validUsers map[string]int) {
	// A valid user exists in both the map and the slice
	// If the user exists, then update their count in the map

	for _, username := range messagedUsers {
		_, ok := validUsers[username]

		if ok {
			validUsers[username] += 1
		}
	}
}

func main() {}
