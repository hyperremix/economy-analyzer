package token

import (
	"net/url"
	"time"

	"golang.org/x/crypto/bcrypt"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/hyperremix/economy-analyzer/backend/api/server"
	"github.com/hyperremix/economy-analyzer/backend/dataAccess/user"
	"github.com/hyperremix/economy-analyzer/backend/model"
)

type TokenController struct {
	server.GetNotSupported
	server.PutNotSupported
	server.DeleteNotSupported
}

const key = "ChangeThis"

func (TokenController) Post(values url.Values) (int, interface{}) {
	user := new(user.UserRepository).Find()

	if err := bcrypt.CompareHashAndPassword(user[0].HashedPassword, []byte(values.Get("password"))); err != nil {
		return 401, ""
	}

	signedToken, err := jwt.New(jwt.SigningMethodHS256).SignedString([]byte(key))

	if err != nil {
		return 500, ""
	}

	token := model.Token{AccessToken: signedToken, CreatedAt: time.Now(), TokenType: "bearer"}

	// TODO insert token in repository

	return 200, token
}
