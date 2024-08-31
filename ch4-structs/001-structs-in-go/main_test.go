package structs

import (
	"fmt"
	"testing"
)

func getMessageText(m messageToSend) string {
	return fmt.Sprintf("Sending message: '%s' to: %v", m.message, m.phoneNumber)
}

func Test(t *testing.T) {
	type testCase struct {
		phoneNumber int
		message     string
		expected    string
	}
	tests := []testCase{
		{148255510981, "Thanks for signing up", "Sending message: 'Thanks for signing up' to: 148255510981"},
		{148255510982, "Love to have you aboard!", "Sending message: 'Love to have you aboard!' to: 148255510982"},
	}
	if withSubmit {
		tests = append(tests, []testCase{
			{148255510983, "We're so excited to have you", "Sending message: 'We're so excited to have you' to: 148255510983"},
			{148255510984, "", "Sending message: '' to: 148255510984"},
			{148255510985, "Hello, World!", "Sending message: 'Hello, World!' to: 148255510985"},
		}...)
	}

	passCount := 0
	failCount := 0

	for _, test := range tests {
		output := getMessageText(messageToSend{
			phoneNumber: test.phoneNumber,
			message:     test.message,
		})
		if output != test.expected {
			failCount++
			t.Errorf(`---------------------------------
Inputs:     (%v, %v)
Expecting:  %v
Actual:     %v
Fail`, test.phoneNumber, test.message, test.expected, output)
		} else {
			passCount++
			fmt.Printf(`---------------------------------
Inputs:     (%v, %v)
Expecting:  %v
Actual:     %v
Pass
`, test.phoneNumber, test.message, test.expected, output)
		}
	}

	fmt.Println("---------------------------------")
	fmt.Printf("%d passed, %d failed\n", passCount, failCount)
}

// withSubmit is set at compile time depending
// on which button is used to run the tests
var withSubmit = true
