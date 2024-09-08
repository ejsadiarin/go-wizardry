package structs

type User struct {
	Name string
	Membership
}

type Membership struct {
	Type             string
	MessageCharLimit int
}

func newUser(name string, membershipType string) User {
	u := User{
		Name: name,
		Membership: Membership{
			Type:             membershipType,
			MessageCharLimit: 100,
		},
	}

	if membershipType == "premium" {
		u.MessageCharLimit = 1000
	}

	return u
}
