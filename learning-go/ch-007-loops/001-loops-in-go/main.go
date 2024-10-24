package loops

func bulkSend(numMessages int) float64 {
	cost := 0.0
	fee := 0.0
	for i := 0; i < numMessages; i++ {
		if i > 0 {
			fee += 0.01
		}
		cost += 1.0 + fee
	}
	return cost
}
