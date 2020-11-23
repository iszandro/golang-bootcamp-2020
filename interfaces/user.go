package interfaces

import "github.com/iszandro/golang-bootcamp-2020/models"

// UserRepository defines the methods user repositories must respond.
type UserRepository interface {
	GetUsers() ([]models.User, error)
}
