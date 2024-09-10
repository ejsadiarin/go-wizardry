package interfaces

import "fmt"

func (e email) cost() int {
	costInCents := len(e.body) * 2
	if !e.isSubscribed {
		costInCents = len(e.body) * 5
	}

	return costInCents
}

func (e email) format() string {
	s := "Subscribed"
	if !e.isSubscribed {
		s = "Not Subscribed"
	}
	return fmt.Sprintf("'%s' | %s", e.body, s)
}

type expense interface {
	cost() int
}

type formatter interface {
	format() string
}

type email struct {
	isSubscribed bool
	body         string
}
