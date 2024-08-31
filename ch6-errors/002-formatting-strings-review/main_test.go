package error

import (
	"fmt"
	"testing"
)

func Test(t *testing.T) {
	type testCase struct {
		cost      float64
		recipient string
		expected  string
	}
	tests := []testCase{
		{1.4, "+1 (435) 555 0923", "SMS that costs $1.40 to be sent to '+1 (435) 555 0923' can not be sent"},
		{2.1, "+2 (702) 555 3452", "SMS that costs $2.10 to be sent to '+2 (702) 555 3452' can not be sent"},
	}
	if withSubmit {
		tests = append(tests, []testCase{
			{32.1, "+1 (801) 555 7456", "SMS that costs $32.10 to be sent to '+1 (801) 555 7456' can not be sent"},
			{14.4, "+1 (234) 555 6545", "SMS that costs $14.40 to be sent to '+1 (234) 555 6545' can not be sent"},
		}...)
	}

	passCount := 0
	failCount := 0

	for _, test := range tests {
		output := getSMSErrorString(test.cost, test.recipient)
		if output != test.expected {
			failCount++
			t.Errorf(`---------------------------------
Inputs:     (%v, %v)
Expecting:  %v
Actual:     %v
Fail`, test.cost, test.recipient, test.expected, output)
		} else {
			passCount++
			fmt.Printf(`---------------------------------
Inputs:     (%v, %v)
Expecting:  %v
Actual:     %v
Pass
`, test.cost, test.recipient, test.expected, output)
		}
	}

	fmt.Println("---------------------------------")
	fmt.Printf("%d passed, %d failed\n", passCount, failCount)
}

// withSubmit is set at compile time depending on which button is used to run the tests
var withSubmit = true
