package maps

func deleteIfNecessary(users map[string]user, name string) (deleted bool, err error) {
	// ?
}

type user struct {
	name                 string
	number               int
	scheduledForDeletion bool
}
