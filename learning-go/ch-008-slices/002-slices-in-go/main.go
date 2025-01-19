package slices

import "fmt"

const (
	planFree = "free"
	planPro  = "pro"
)

func getMessageWithRetriesForPlan(plan string, messages [3]string) ([]string, error) {
	if plan == planPro {
		return messages[:], nil
	}

	if plan == planFree {
		return messages[:2], nil
	}

	return nil, fmt.Errorf("unsupported plan")
}
