package monthly

import (
	"github.com/gin-gonic/gin"
	"github.com/hyperremix/economy-analyzer/backend/application"
	"net/http"
)

type MonthlyController struct {
	monthlyFacade *application.MonthlyFacade
}

const path = "/monthlies"

func RegisterMonthlyController(router *gin.Engine, routePrefix string) {
	mc := &MonthlyController{monthlyFacade: application.NewMonthlyFacade()}

	router.GET(routePrefix+path, mc.Get())
}

func (mc *MonthlyController) Get() gin.HandlerFunc {
	return func(c *gin.Context) {
		monthlies := mc.monthlyFacade.Find()

		c.JSON(http.StatusOK, ManyNewMonthlyApiModels(monthlies))
	}
}
