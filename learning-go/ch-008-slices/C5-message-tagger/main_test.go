package slices

import (
	"fmt"
	"reflect"
	"testing"
)

func TestTagMessages(t *testing.T) {
	tests := []struct {
		messages []sms
		expected [][]string
	}{
		{
			messages: []sms{{id: "001", content: "Urgent, please respond!"}, {id: "002", content: "Big sale on all items!"}},
			expected: [][]string{{"Urgent"}, {"Promo"}},
		},
		{
			messages: []sms{{id: "003", content: "Enjoy your day"}},
			expected: [][]string{{}},
		},
	}

	if withSubmit {
		tests = append(tests, struct {
			messages []sms
			expected [][]string
		}{
			messages: []sms{{id: "004", content: "Sale! Don't miss out on these urgent promotions!"}},
			expected: [][]string{{"Urgent", "Promo"}},
		})
	}

	passCount := 0
	failCount := 0

	for _, test := range tests {
		actual := tagMessages(test.messages, tagger)
		for i, msg := range actual {
			if !reflect.DeepEqual(msg.tags, test.expected[i]) {
				failCount++
				t.Errorf(`---------------------------------
Test Failed for message ID %s
Expecting: %v
Actual:    %v
Fail`, msg.id, test.expected[i], msg.tags)
			} else {
				passCount++
				fmt.Printf(`---------------------------------
Test Passed for message ID %s
Expecting: %v
Actual:    %v
Pass
`, msg.id, test.expected[i], msg.tags)
			}
		}
	}

	fmt.Println("---------------------------------")
	fmt.Printf("%d passed, %d failed\n", passCount, failCount)
}

// withSubmit is set at compile time depending on which button is used to run the tests
var withSubmit = true
