package repositories

import "github.com/iszandro/golang-bootcamp-2020/models"

// MemoryUserRepository allows to get users from memory.
// In this case, from users variable.
type MemoryUserRepository struct{}

var users = []models.User{
	{ID: 1, FirstName: "Harry", LastName: "Potter"},
	{ID: 2, FirstName: "Albus", LastName: "Dumbledore"},
	{ID: 3, FirstName: "Tom", LastName: "Riddle"},
}

// GetUsers return a slice of models.User by reading a users variable stored in memory.
func (repository MemoryUserRepository) GetUsers() ([]models.User, error) {
	return users, nil
}
