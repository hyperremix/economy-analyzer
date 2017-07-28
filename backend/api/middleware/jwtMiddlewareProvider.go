package middleware

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/hyperremix/economy-analyzer/backend/dataAccess"
	"github.com/hyperremix/economy-analyzer/backend/model"
	"net/http"
	"strings"
	"time"
)

type JWTMiddlewareProvider struct {
	tokenRepository *dataAccess.TokenRepository
}

const AuthTokenKey = "token"

func NewJWTMiddlewareProvider() *JWTMiddlewareProvider {
	return &JWTMiddlewareProvider{tokenRepository: dataAccess.NewTokenRepository()}
}

func (jp *JWTMiddlewareProvider) Get() gin.HandlerFunc {
	return func(c *gin.Context) {
		value := c.GetHeader("Authorization")

		tokenString, err := jp.parseBearerAuthHeader(value)
		if err != nil {
			c.AbortWithError(http.StatusUnauthorized, err)
		}

		token, err := jp.tokenRepository.FindSingleByAccessToken(tokenString)
		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		if err := jp.validateToken(token); err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		c.Set(AuthTokenKey, token)
	}
}

func (jp *JWTMiddlewareProvider) parseBearerAuthHeader(value string) (tokenString string, err error) {
	const headerSeparator = " "
	const expectedParts = 2

	parts := strings.Split(value, headerSeparator)

	if len(parts) != expectedParts || parts[0] != "Bearer" {
		err = errors.New("authorization failed")
		return
	}

	tokenString = parts[1]
	return
}

func (jp *JWTMiddlewareProvider) validateToken(token model.Token) (err error) {
	expiresAt := token.CreatedAt.Add(time.Second * time.Duration(token.ExpiresIn))
	if expiresAt.Before(time.Now()) {
		err = errors.New("token expired")
	}

	return
}
