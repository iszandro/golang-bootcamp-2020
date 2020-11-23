package interactors

import (
	"github.com/iszandro/golang-bootcamp-2020/interfaces"
	"github.com/iszandro/golang-bootcamp-2020/models"
)

// UserInteractor has all user use cases. In this case, the only
// use case is to get a list of users.
// It uses a interfaces.UserRepository to get all users by calling
// ListAllUsers().
type UserInteractor struct {
	repository interfaces.UserRepository
}

func NewUserInteractor(repository interfaces.UserRepository) *UserInteractor {
	return &UserInteractor{
		repository: repository,
	}
}

// ListAllUsers returns a slice of models.User.
func (interactor *UserInteractor) ListAllUsers() ([]models.User, error) {
	return interactor.repository.GetUsers()
}

func (interactor *UserInteractor) SetRepository(repository interfaces.UserRepository) {
	interactor.repository = repository
}
