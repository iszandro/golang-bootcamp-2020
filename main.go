package main

import (
	"fmt"
	"log"

	"github.com/iszandro/golang-bootcamp-2020/adapters"
	"github.com/iszandro/golang-bootcamp-2020/interactors"
	"github.com/iszandro/golang-bootcamp-2020/models"
	"github.com/iszandro/golang-bootcamp-2020/repositories"
)

func main() {
	fmt.Println("**********", "CSV Users", "**********")
	interactor := interactors.NewUserInteractor(adapters.NewCSVUserRepositoryAdapter())
	users, err := interactor.ListAllUsers()

	if err != nil {
		log.Fatal(err)
	}

	printUsers(users)

	fmt.Println("**********", "Memory Users", "**********")
	interactor.SetRepository(repositories.MemoryUserRepository{})
	users, err = interactor.ListAllUsers()

	if err != nil {
		log.Fatal(err)
	}

	printUsers(users)
}

func printUsers(users []models.User) {
	for _, user := range users {
		fmt.Printf("User #%d\nFirst name: %s\nLast name: %s\nEmail: %s\n\n",
			user.ID, user.FirstName, user.LastName, user.Email())
	}
}
