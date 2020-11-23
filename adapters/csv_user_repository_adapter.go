package adapters

import (
	"github.com/iszandro/golang-bootcamp-2020/models"
	"github.com/iszandro/golang-bootcamp-2020/repositories"
)

// CSVUserRepositoryAdapter allows to adapt repositories.CVSUserRepository into a
// interfaces.UserRepository interface.
// Due to repositories.CVSUserRepository needs a file path for its GetUsers(filePath string) method,
// this adapter responds to GetUsers() method by setting file path.
// If FilePath is empty, defaultUsersPath constant is used.
type CSVUserRepositoryAdapter struct {
	csvUserRepository repositories.CVSUserRepository
	FilePath          string
}

const defaultUsersPath = "fixtures/files/users.csv"

func NewCSVUserRepositoryAdapter() CSVUserRepositoryAdapter {
	return CSVUserRepositoryAdapter{
		csvUserRepository: repositories.CVSUserRepository{},
	}
}

// GetUsers is an adapted method which returns a slice of models.User.
func (adapter CSVUserRepositoryAdapter) GetUsers() ([]models.User, error) {
	var path string

	if len(adapter.FilePath) > 0 {
		path = adapter.FilePath
	} else {
		path = defaultUsersPath
	}

	return adapter.csvUserRepository.GetUsers(path)
}
