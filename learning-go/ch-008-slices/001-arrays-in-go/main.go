package slices

func getMessageWithRetries(primary, secondary, tertiary string) ([3]string, [3]int) {
	f := len(primary)
	s := f + len(secondary)
	t := s + len(tertiary)
	return [3]string{primary, secondary, tertiary}, [3]int{f, s, t}
}
