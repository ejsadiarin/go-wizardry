package loops

func maxMessages(thresh int) int {
	messageCost := 0
	var i int
	for i = 0; messageCost <= thresh; i++ {
		messageCost += 100 + i
	}
	return i - 1
}
