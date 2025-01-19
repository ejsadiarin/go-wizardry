package slices

func getMessageCosts(messages []string) []float64 {
	msgCosts := make([]float64, len(messages))

	for i, m := range messages {
		cost := float64(len(m)) * 0.01
		msgCosts[i] = cost
	}

	return msgCosts
}
