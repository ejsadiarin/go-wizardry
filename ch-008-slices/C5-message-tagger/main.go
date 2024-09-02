package slices

type sms struct {
	id      string
	content string
	tags    []string
}

func tagMessages(messages []sms, tagger func(sms) []string) []sms {
	// ?
}

func tagger(msg sms) []string {
	tags := []string{}
	// ?
}
