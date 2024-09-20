package error

import (
	"fmt"
)

func sendSMSToCouple(msgToCustomer, msgToSpouse string) (int, error) {
	mCustomer, err := sendSMS(msgToCustomer)
	if err != nil {
		return 0, err
	}

	mSpouse, err := sendSMS(msgToSpouse)
	if err != nil {
		return 0, err
	}

	return (mCustomer + mSpouse), nil
}

// don't edit below this line

func sendSMS(message string) (int, error) {
	const maxTextLen = 25
	const costPerChar = 2
	if len(message) > maxTextLen {
		return 0, fmt.Errorf("can't send texts over %v characters", maxTextLen)
	}
	return costPerChar * len(message), nil
}
