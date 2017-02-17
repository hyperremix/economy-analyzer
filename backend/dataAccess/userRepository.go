package dataAccess

import (
	"encoding/csv"
	"os"

	"log"

	"github.com/hyperremix/economy-analyzer/backend/model"
)

type UserRepository struct {
	userMap *userMap
}

func NewUserRepository() *UserRepository {
	return &UserRepository{userMap: new(userMap)}
}

func (UserRepository *UserRepository) Find() []model.User {
	fileReader, err := os.Open("C:\\Users\\fredr_000\\goplayground\\src\\github.com\\hyperremix\\economy-analyzer\\user.csv")

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

	return UserRepository.userMap.TransformMany(records)
}
