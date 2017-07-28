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
	jwtAuthMiddlewareProvider   *middleware.JWTMiddlewareProvider
}

const (
	tokenPath        = "/auth/token"
	refreshTokenPath = "/auth/refresh_token"
)

func RegisterTokenController(router *gin.Engine, routePrefix string) {
	tc := &tokenController{
		userRepository:              dataAccess.NewUserRepository(),
		tokenRepository:             dataAccess.NewTokenRepository(),
		basicAuthMiddlewareProvider: middleware.NewBasicAuthMiddlewareProvider(),
		jwtAuthMiddlewareProvider:   middleware.NewJWTMiddlewareProvider()}

	tokenEndpoint := router.Group(routePrefix+tokenPath, tc.basicAuthMiddlewareProvider.Get())
	tokenEndpoint.POST("", tc.PostToken())

	refreshTokenEndpoint := router.Group(routePrefix+refreshTokenPath, tc.jwtAuthMiddlewareProvider.Get())
	refreshTokenEndpoint.PUT("", tc.PutRefreshToken())
}

func (tc *tokenController) PostToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		tc.createToken(c)
	}
}

func (tc *tokenController) PutRefreshToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.MustGet(middleware.AuthTokenKey).(model.Token)
		token.Refresh()

		tc.tokenRepository.Upsert(token)

		c.JSON(http.StatusOK, NewTokenApiModel(token))
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
