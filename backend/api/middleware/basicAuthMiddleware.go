package middleware

import (
	"encoding/base64"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/hyperremix/economy-analyzer/backend/dataAccess"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2/bson"
	"net/http"
	"strings"
)

type basicAuthMiddleWareProvider struct {
	userRepository *dataAccess.UserRepository
}

const AuthUserIDKey = "userID"

func NewBasicAuthMiddleWareProvider() *basicAuthMiddleWareProvider {
	return &basicAuthMiddleWareProvider{userRepository: dataAccess.NewUserRepository()}
}

func (bm *basicAuthMiddleWareProvider) Get() gin.HandlerFunc {
	return func(c *gin.Context) {
		value := c.GetHeader("Authorization")

		username, password, err := bm.parseAuthorizationHeader(value)
		if err != nil {
			c.AbortWithError(http.StatusUnauthorized, err)
		}

		userID, err := bm.validateUser(username, password)
		if err != nil {
			c.AbortWithError(http.StatusUnauthorized, err)
		}

		c.Set(AuthUserIDKey, userID)
	}
}

func (bm *basicAuthMiddleWareProvider) parseAuthorizationHeader(value string) (username, password string, err error) {
	const headerSeparator = " "
	const valueSeparator = ":"
	const expectedParts = 2

	parts := strings.Split(value, headerSeparator)

	if len(parts) != expectedParts || parts[0] != "Basic" {
		err = errors.New("authorization failed")
		return
	}

	payload, _ := base64.StdEncoding.DecodeString(parts[1])
	pair := strings.Split(string(payload), valueSeparator)

	if len(pair) != expectedParts {
		err = errors.New("authorization failed")
		return
	}

	username = pair[0]
	password = pair[1]
	return
}

func (bm *basicAuthMiddleWareProvider) validateUser(username string, password string) (userID bson.ObjectId, err error) {
	user, err := bm.userRepository.FindSingleByUsername(username)
	if err != nil {
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.HashedPassword), []byte(password))
	if err != nil {
		return
	}

	userID = user.ID
	return
}
