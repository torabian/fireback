package fireback

import (
	"os/user"
)

// Safer option, if you are running go in the phone.

func GetOsUserWithPhone() *user.User {

	currentUser, err := user.Current()
	if err != nil {
		return &user.User{
			Uid:      "21",
			Gid:      "21",
			Username: "local",
			Name:     "local",
			HomeDir:  "local",
		}
	}

	return currentUser

}
