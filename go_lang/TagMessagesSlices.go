package main

import (
	"strings"
)

/*
Complete the tagMessages function.
It should take a slice of sms messages, and a function (that takes a sms as input and returns a slice of strings) as inputs.
And it should return a slice of sms messages.

It should loop through each message and set the tags to the result of the passed in function.
Be sure to modify the messages of the original slice using bracket notation messages[i].

Complete the tagger function. It should take a sms message and return a slice of strings.
Return an initialized slice, even if no tags match. No nil slices.
For any message that contains "urgent" (regardless of casing) in the content, the Urgent tag should be applied first.
For any message that contains "sale" (regardless of casing), the Promo tag should be applied second.
*/

type sms struct {
	id      string
	content string
	tags    []string
}

func tagMessages(messages []sms, tagger func(sms) []string) []sms {
	for idx, message := range messages {
		messages[idx] = sms{
			id:      message.id,
			content: message.content,
			tags:    tagger(message),
		}
	}
	return messages
}

func tagger(message sms) []string {
	tags := []string{}
	hasUrgent := strings.Contains(strings.ToLower(message.content), "urgent")
	hasSale := strings.Contains(strings.ToLower(message.content), "sale")

	if hasUrgent {
		tags = append(tags, "Urgent")
	}

	if hasSale {
		tags = append(tags, "Promo")
	}
	return tags

}

func main() {
	//	CMD to run file against TestCases:
	//	go test -v TagMessagesSlices TagMessagesSlices_test.go
}
