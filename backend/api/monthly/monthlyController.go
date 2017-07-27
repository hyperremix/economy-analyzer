package monthly

import (
	"github.com/gin-gonic/gin"
	"github.com/hyperremix/economy-analyzer/backend/api/middleware"
	"github.com/hyperremix/economy-analyzer/backend/application"
	"net/http"
)

type MonthlyController struct {
	monthlyFacade         *application.MonthlyFacade
	jwtMiddlewareProvider *middleware.JWTMiddlewareProvider
}

const path = "/monthlies"

func RegisterMonthlyController(router *gin.Engine, routePrefix string) {
	mc := &MonthlyController{
		monthlyFacade:         application.NewMonthlyFacade(),
		jwtMiddlewareProvider: middleware.NewJWTMiddlewareProvider()}

	endpoint := router.Group(routePrefix+path, mc.jwtMiddlewareProvider.Get())
	endpoint.GET("", mc.Get())
}

func (mc *MonthlyController) Get() gin.HandlerFunc {
	return func(c *gin.Context) {
		monthlies := mc.monthlyFacade.Find()

		c.JSON(http.StatusOK, ManyNewMonthlyApiModels(monthlies))
	}
}
