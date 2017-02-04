package version

import (
	"net/url"

	"github.com/hyperremix/economy-analyzer/api"
)

//VersionController responsible for the version resource
type VersionController struct {
	api.PostNotSupported
	api.PutNotSupported
	api.DeleteNotSupported
}

//Get returns the current version of the application
func (VersionController) Get(values url.Values) (int, interface{}) {
	data := map[string]string{"version": "0.1.0"}
	return 200, data
}
