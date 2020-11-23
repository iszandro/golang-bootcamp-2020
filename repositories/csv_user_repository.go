package repositories

import (
	"encoding/csv"
	"io"
	"os"

	"github.com/iszandro/golang-bootcamp-2020/models"
)

// CVSUserRepository allows to get users from a CSV file.
type CVSUserRepository struct{}

// GetUsers returns a slice of models.User by reading a CSV file.
// It ignores first row (index 0), which has the CSV headers.
// It loops each CSV row and builds a models.User with the CSV data.
func (repository CVSUserRepository) GetUsers(filePath string) ([]models.User, error) {
	file, err := os.Open(filePath)
	defer file.Close()
	if err != nil {
		return nil, err
	}

	var users []models.User
	reader := csv.NewReader(file)

	for i := 0; ; i++ {
		record, err := reader.Read()

		if i == 0 {
			continue
		}

		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}

		users = append(users, buildUserFromRow(i, record))
	}

	return users, nil
}

func buildUserFromRow(id int, row []string) models.User {
	user := models.User{ID: id, FirstName: row[0], LastName: row[1]}
	user.SetEmail(row[2])
	return user
}
