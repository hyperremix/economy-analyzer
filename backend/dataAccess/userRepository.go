package dataAccess

import (
	"github.com/hyperremix/economy-analyzer/backend/model"
	"gopkg.in/mgo.v2/bson"
)

type UserRepository struct{}

func NewUserRepository() *UserRepository {
	return new(UserRepository)
}

func (ur *UserRepository) FindSingleByUsername(username string) (result model.User, err error) {
	err = findSingle("users", bson.M{"username": username}, &result)
	return
}

func (ur *UserRepository) Insert(user model.User) (err error) {
	err = insert("users", user)
	return
}
