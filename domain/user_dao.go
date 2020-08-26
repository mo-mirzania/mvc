package domain

import (
	"github.com/mo-mirzania/test/utils/resterror"
)

var usersDB = map[int]*User{
	1: {
		ID:        1,
		FirstName: "Mohsen",
		LastName:  "Mirzania",
		Email:     "mo.mirzania@gmail.com",
	},
}

// GetUser method
func (u *User) GetUser() (*User, *resterror.RestErr) {
	user := usersDB[u.ID]
	if user == nil {
		return nil, resterror.NotFoundError("User not found in the database")
	}
	return user, nil
}
