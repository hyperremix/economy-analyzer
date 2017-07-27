package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/hyperremix/economy-analyzer/backend/api/monthly"
	"github.com/hyperremix/economy-analyzer/backend/api/registration"
	"github.com/hyperremix/economy-analyzer/backend/api/token"
	"github.com/hyperremix/economy-analyzer/backend/api/version"
)

type API struct {
	router *gin.Engine
}

const routePrefix = "/api"

func NewAPI(router *gin.Engine) *API {
	return &API{router: router}
}

func (api *API) RegisterControllers() {
	monthly.RegisterMonthlyController(api.router, routePrefix)
	registration.RegisterRegistrationController(api.router, routePrefix)
	token.RegisterTokenController(api.router, routePrefix)
	version.RegisterVersionController(api.router, routePrefix)
}

func (api *API) Start(port int) {
	portString := fmt.Sprintf(":%d", port)
	api.router.Run(portString)
}
