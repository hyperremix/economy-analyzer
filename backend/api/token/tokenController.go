package token

import (
	"encoding/base64"
	"errors"
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/hyperremix/economy-analyzer/backend/dataAccess"
	"github.com/hyperremix/economy-analyzer/backend/model"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/mgo.v2/bson"
	"net/http"
	"strings"
	"time"
)

// TODO: Change this
var key = []byte("ChangeThis")

type tokenController struct {
	userRepository  *dataAccess.UserRepository
	tokenRepository *dataAccess.TokenRepository
}

const (
	path        = "/token"
	authUserKey = "user"
)

func RegisterTokenController(router *gin.Engine, routePrefix string) {
	tc := &tokenController{
		userRepository:  dataAccess.NewUserRepository(),
		tokenRepository: dataAccess.NewTokenRepository()}

	endpoint := router.Group(routePrefix+path, tc.basicAuthMiddleware())
	endpoint.POST("", tc.Post())
}

func (tc *tokenController) Post() gin.HandlerFunc {
	return func(c *gin.Context) {
		tc.createToken(c)
	}
}

func (tc *tokenController) createToken(c *gin.Context) {
	signedToken, err := jwt.New(jwt.SigningMethodHS256).SignedString([]byte(key))

	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
	}

	userID := c.MustGet(authUserKey).(bson.ObjectId)
	token := model.Token{
		UserID:      userID,
		AccessToken: signedToken,
		CreatedAt:   time.Now(),
		ExpiresIn:   (time.Hour * 24).Seconds()}

	tc.tokenRepository.Upsert(token)

	c.JSON(http.StatusCreated, NewTokenApiModel(token))
}

func (tc *tokenController) basicAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		value := c.GetHeader("Authorization")

		username, password, err := tc.parseAuthorizationHeader(value)
		if err != nil {
			c.AbortWithError(http.StatusUnauthorized, err)
		}

		userID, err := tc.validateUser(username, password)
		if err != nil {
			c.AbortWithError(http.StatusUnauthorized, err)
		}

		c.Set(authUserKey, userID)
	}
}

func (tc *tokenController) parseAuthorizationHeader(value string) (username, password string, err error) {
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

func (tc *tokenController) validateUser(username string, password string) (userID bson.ObjectId, err error) {
	user, err := tc.userRepository.FindSingleByUsername(username)
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
