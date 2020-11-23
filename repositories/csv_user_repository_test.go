package repositories

import (
	"testing"

	"github.com/iszandro/golang-bootcamp-2020/models"
)

func TestCVSUserRepository_GetUsers_NonExistingFile(t *testing.T) {
	repository := CVSUserRepository{}
	_, err := repository.GetUsers("not/a/file.csv")

	if err == nil {
		t.Fatal("File should not exist")
	}

	if exp := "open not/a/file.csv: no such file or directory"; exp != err.Error() {
		t.Fatalf("expected \"%s\", actual \"%s\"", exp, err.Error())
	}
}

func TestCVSUserRepository_GetUsers(t *testing.T) {
	repository := CVSUserRepository{}
	users, _ := repository.GetUsers("../fixtures/files/users.csv")
	expUsers := []models.User{
		{ID: 1, FirstName: "Bruce", LastName: "Wayne"},
		{ID: 2, FirstName: "Clark", LastName: "Kent"},
		{ID: 3, FirstName: "Diana", LastName: "Prince"},
	}
	expUsers[0].SetEmail("bruce@wayne.com")
	expUsers[1].SetEmail("clark@kent.com")
	expUsers[2].SetEmail("diana@prince.com")

	if len(expUsers) != len(users) {
		t.Fatalf("expected %d, actual %d", len(expUsers), len(users))
	}

	for i := 0; i < len(users); i++ {
		if expUsers[i] != users[i] {
			t.Fatalf("expected \"%v\", actual \"%v\"", expUsers[i], users[i])
		}
	}
}
