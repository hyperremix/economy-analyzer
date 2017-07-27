package token

import (
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/hyperremix/economy-analyzer/backend/api/server"
	"github.com/hyperremix/economy-analyzer/backend/dataAccess"
	"github.com/hyperremix/economy-analyzer/backend/model"
	"golang.org/x/crypto/bcrypt"
	"net/url"
	"time"
)

// TODO: Change this
var key = []byte("ChangeThis")

type tokenController struct {
	server.GetNotSupported
	server.PutNotSupported
	server.DeleteNotSupported
	userRepository  *dataAccess.UserRepository
	tokenRepository *dataAccess.TokenRepository
}

func NewTokenController() *tokenController {
	return &tokenController{
		userRepository:  dataAccess.NewUserRepository(),
		tokenRepository: dataAccess.NewTokenRepository()}
}

func (tc *tokenController) Post(values url.Values) (int, interface{}) {
	username := values.Get("client_id")
	password := []byte(values.Get("client_secret"))

	user, err := tc.userRepository.FindSingleByUsername(username)

	if err != nil {
		return 401, ""
	}

	if err := bcrypt.CompareHashAndPassword(user.HashedPassword, password); err != nil {
		return 401, ""
	}

	signedToken, err := jwt.New(jwt.SigningMethodHS256).SignedString([]byte(key))

	if err != nil {
		return 500, ""
	}

	token := model.Token{
		AccessToken: signedToken,
		CreatedAt:   time.Now(),
		ExpiresIn:   (time.Hour * 24).Seconds()}

	tc.tokenRepository.Insert(token)

	return 200, NewTokenApiModel(token)
}
