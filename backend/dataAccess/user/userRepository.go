package user

import (
	"encoding/csv"
	"os"

	"log"

	"github.com/hyperremix/economy-analyzer/backend/model"
)

const repositoryPath string = "C:\\Users\\fredr_000\\goplayground\\src\\github.com\\hyperremix\\economy-analyzer\\user.csv"

type UserRepository struct{}

func (userRepository *UserRepository) Find() []model.User {
	fileReader, err := os.Open(repositoryPath)

	if err != nil {
		log.Fatal(err)
		return make([]model.User, 0)
	}

	reader := csv.NewReader(fileReader)
	reader.Comma = ';'

	reader.Read()

	records, err := reader.ReadAll()

	if err != nil {
		log.Fatal(err)
		return make([]model.User, 0)
	}

	return new(UserMap).TransformMany(records)
}
