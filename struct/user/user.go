package user

import (
	"errors"
	"time"
)

type User struct {
	firstName string
	lastName  string
	dob       string
	createdAt time.Time
}

func New(firstName, lastName, dob string) (*User, error) {
	if firstName == "" || lastName == "" || dob == "" {
		return nil, errors.New("all fields are required")
	}

	return &User{
		firstName: firstName,
		lastName:  lastName,
		dob:       dob,
		createdAt: time.Now(),
	}, nil
}

func (u *User) ClearUsername() {
	u.firstName = ""
	u.lastName = ""
}
