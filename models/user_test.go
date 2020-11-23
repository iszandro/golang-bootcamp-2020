package models

import (
	"fmt"
	"testing"
)

func TestUser(t *testing.T) {
	user := User{ID: 1, FirstName: "Bruce", LastName: "Wayne"}
	user.SetEmail("bruce@wayne.com")

	exp := "ID: 1, FirstName: Bruce, LastName: Wayne, Email: bruce@wayne.com"
	act := fmt.Sprintf("ID: %d, FirstName: %s, LastName: %s, Email: %s",
		user.ID, user.FirstName, user.LastName, user.Email())

	if exp != act {
		t.Fatalf("expected \"%s\", actual \"%s\"", exp, act)
	}
}

func TestUser_SetEmail(t *testing.T) {
	user := User{}

	user.SetEmail("not email")
	if user.email != "" {
		t.Fatalf("expected \"%s\", actual \"%s\"", "", user.email)
	}

	validEmail := "bruce@wayne.com"
	user.SetEmail(validEmail)

	if user.email != validEmail {
		t.Fatalf("expected \"%s\", actual \"%s\"", validEmail, user.email)
	}
}
