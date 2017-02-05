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
}

func (MonthlyController) Get(values url.Values) (int, interface{}) {
	return 200, new(application.MonthlyFacade).Find()
}
