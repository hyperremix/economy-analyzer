package user

import "github.com/hyperremix/economy-analyzer/backend/model"

type UserMap struct{}

func (userMap *UserMap) TransformMany(records [][]string) []model.User {

	var users = make([]model.User, len(records))

	for i, record := range records {
		users[i] = userMap.Transform(record)
	}

	return users
}

func (userMap *UserMap) Transform(record []string) model.User {
	return model.User{Username: record[0], HashedPassword: []byte(record[1])}
}
