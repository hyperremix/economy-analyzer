package registration

import (
	"github.com/hyperremix/economy-analyzer/backend/api/server"
	"github.com/hyperremix/economy-analyzer/backend/dataAccess"
	"github.com/hyperremix/economy-analyzer/backend/model"
	"golang.org/x/crypto/bcrypt"
	"net/url"
)

type RegistrationController struct {
	server.GetNotSupported
	server.PutNotSupported
	server.DeleteNotSupported
	userRepository *dataAccess.UserRepository
}

func NewRegistrationController() *RegistrationController {
	return &RegistrationController{userRepository: dataAccess.NewUserRepository()}
}

func (rc *RegistrationController) Post(values url.Values) (int, interface{}) {
	username := values.Get("username")
	password := []byte(values.Get("password"))

	hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)

	if err != nil {
		return 500, ""
	}

	user := model.User{Username: username, HashedPassword: hashedPassword}
	if err := rc.userRepository.Insert(user); err != nil {
		return 500, ""
	}

	return 201, nil
}
