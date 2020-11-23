package repositories

import (
	"testing"

	"github.com/iszandro/golang-bootcamp-2020/models"
)

func TestMemoryUserRepository_GetUsers(t *testing.T) {
	repository := MemoryUserRepository{}
	users, err := repository.GetUsers()

	if err != nil {
		t.Fatalf("expected err to be nil, but was %#v", err)
	}

	expUsers := []models.User{
		{ID: 1, FirstName: "Harry", LastName: "Potter"},
		{ID: 2, FirstName: "Albus", LastName: "Dumbledore"},
		{ID: 3, FirstName: "Tom", LastName: "Riddle"},
	}

	if len(expUsers) != len(users) {
		t.Fatalf("expected %d, actual %d", len(expUsers), len(users))
	}

	for i := 0; i < len(users); i++ {
		if expUsers[i] != users[i] {
			t.Fatalf("expected \"%v\", actual \"%v\"", expUsers[i], users[i])
		}
	}
}
