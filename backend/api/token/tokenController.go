package token

import (
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/hyperremix/economy-analyzer/backend/dataAccess"
	"github.com/hyperremix/economy-analyzer/backend/model"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"time"
)

// TODO: Change this
var key = []byte("ChangeThis")

type tokenController struct {
	userRepository  *dataAccess.UserRepository
	tokenRepository *dataAccess.TokenRepository
}

const path = "/token"

func RegisterTokenController(router *gin.Engine, routePrefix string) {
	tc := &tokenController{
		userRepository:  dataAccess.NewUserRepository(),
		tokenRepository: dataAccess.NewTokenRepository()}
	router.POST(routePrefix+path, tc.Post())
}

func (tc *tokenController) Post() gin.HandlerFunc {
	return func(c *gin.Context) {
		var tokenRequest tokenRequest
		if err := c.BindJSON(&tokenRequest); err == nil {
			tc.createToken(c, tokenRequest)
		}
	}
}
func (tc *tokenController) createToken(c *gin.Context, tokenRequest tokenRequest) {
	user, err := tc.userRepository.FindSingleByUsername(tokenRequest.ClientId)

	if err != nil {
		c.JSON(http.StatusUnauthorized, nil)
		return
	}

	if err := bcrypt.CompareHashAndPassword(user.HashedPassword, []byte(tokenRequest.ClientSecret)); err != nil {
		c.JSON(http.StatusUnauthorized, nil)
		return
	}

	signedToken, err := jwt.New(jwt.SigningMethodHS256).SignedString([]byte(key))

	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	token := model.Token{
		AccessToken: signedToken,
		CreatedAt:   time.Now(),
		ExpiresIn:   (time.Hour * 24).Seconds()}

	tc.tokenRepository.Insert(token)

	c.JSON(http.StatusCreated, NewTokenApiModel(token))
}
