package registration

import (
	"github.com/gin-gonic/gin"
	"github.com/hyperremix/economy-analyzer/backend/dataAccess"
	"github.com/hyperremix/economy-analyzer/backend/model"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

type RegistrationController struct {
	userRepository *dataAccess.UserRepository
}

const path = "/registration"

func RegisterRegistrationController(router *gin.Engine, routePrefix string) {
	rc := &RegistrationController{userRepository: dataAccess.NewUserRepository()}
	router.POST(routePrefix+path, rc.Post())
	return
}

func (rc *RegistrationController) Post() gin.HandlerFunc {
	return func(c *gin.Context) {
		var registration registrationRequest
		if err := c.BindJSON(&registration); err == nil {
			rc.registerUser(c, registration)
		}
	}
}

func (rc *RegistrationController) registerUser(c *gin.Context, request registrationRequest) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)

	if err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	user := model.User{Username: request.Username, HashedPassword: string(hashedPassword)}
	if err := rc.userRepository.Insert(user); err != nil {
		c.JSON(http.StatusInternalServerError, nil)
		return
	}

	c.JSON(http.StatusCreated, nil)
}
