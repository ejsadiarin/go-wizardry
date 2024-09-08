package structs

import "fmt"

type authenticationInfo struct {
	username string
	password string
}

// create the method below

func (a authenticationInfo) getBasicAuth() string {
	return fmt.Sprintf("Authorization: Basic %s:%s", a.username, a.password)
}
