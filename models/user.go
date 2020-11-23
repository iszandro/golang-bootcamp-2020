package models

import "regexp"

// User is a representation of a user.
type User struct {
	ID        int
	FirstName string
	LastName  string
	email     string
}

const emailRegexp = `\A[^@\s]+@[^@\s]+\z`

// SetEmail sets a valid email by using emailRegexp to check if
// email is valid or not.
func (u *User) SetEmail(email string) {
	if matched, _ := regexp.MatchString(emailRegexp, email); matched {
		u.email = email
	}
}

func (u *User) Email() string {
	return u.email
}
