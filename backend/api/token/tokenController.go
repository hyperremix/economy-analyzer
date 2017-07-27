package token

import (
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/hyperremix/economy-analyzer/backend/api/middleware"
	"github.com/hyperremix/economy-analyzer/backend/dataAccess"
	"github.com/hyperremix/economy-analyzer/backend/model"
	"gopkg.in/mgo.v2/bson"
	"net/http"
	"time"
)

// TODO: Change this
var key = []byte("ChangeThis")

type tokenController struct {
	userRepository              *dataAccess.UserRepository
	tokenRepository             *dataAccess.TokenRepository
	basicAuthMiddlewareProvider *middleware.BasicAuthMiddlewareProvider
}

const path = "/token"

func RegisterTokenController(router *gin.Engine, routePrefix string) {
	tc := &tokenController{
		userRepository:              dataAccess.NewUserRepository(),
		tokenRepository:             dataAccess.NewTokenRepository(),
		basicAuthMiddlewareProvider: middleware.NewBasicAuthMiddlewareProvider()}

	endpoint := router.Group(routePrefix+path, tc.basicAuthMiddlewareProvider.Get())
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

	userID := c.MustGet(middleware.AuthUserIDKey).(bson.ObjectId)
	token := model.Token{
		UserID:      userID,
		AccessToken: signedToken,
		CreatedAt:   time.Now(),
		ExpiresIn:   (time.Hour * 24).Seconds()}

	tc.tokenRepository.Upsert(token)

	c.JSON(http.StatusCreated, NewTokenApiModel(token))
}
