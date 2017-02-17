package dataAccess

import "github.com/hyperremix/economy-analyzer/backend/model"

type UserRepository struct{}

func NewUserRepository() *UserRepository {
	return new(UserRepository)
}

func (UserRepository *UserRepository) FindMany() []model.User {
	var results []model.User

	findMany("users", &results)
	return results
}
