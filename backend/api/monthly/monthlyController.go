package monthly

import (
	"net/url"

	"github.com/hyperremix/economy-analyzer/backend/api/server"
	"github.com/hyperremix/economy-analyzer/backend/application"
)

type MonthlyController struct {
	server.PostNotSupported
	server.PutNotSupported
	server.DeleteNotSupported
	monthlyFacade *application.MonthlyFacade
}

func NewMonthlyController() *MonthlyController {
	return &MonthlyController{monthlyFacade: application.NewMonthlyFacade()}
}

func (monthlyController *MonthlyController) Get(values url.Values) (int, interface{}) {
	return 200, new(application.MonthlyFacade).Find()
}
