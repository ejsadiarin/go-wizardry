package loops

func countConnections(groupSize int) int {
	var c int
	for i := 0; i < groupSize; i++ {
		c += i
	}

	return c
}
