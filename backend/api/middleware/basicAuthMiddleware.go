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

const AuthUserIDKey = "userID"

func BasicAuthMiddleWare(userRepository *dataAccess.UserRepository) gin.HandlerFunc {
	return func(c *gin.Context) {
		value := c.GetHeader("Authorization")

		username, password, err := parseAuthorizationHeader(value)
		if err != nil {
			c.AbortWithError(http.StatusUnauthorized, err)
		}

		userID, err := validateUser(userRepository, username, password)
		if err != nil {
			c.AbortWithError(http.StatusUnauthorized, err)
		}

		c.Set(AuthUserIDKey, userID)
	}
}

func parseAuthorizationHeader(value string) (username, password string, err error) {
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

func validateUser(userRepository *dataAccess.UserRepository, username string, password string) (userID bson.ObjectId, err error) {
	user, err := userRepository.FindSingleByUsername(username)
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
