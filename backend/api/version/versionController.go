package version

import (
	"net/url"

	"github.com/hyperremix/economy-analyzer/backend/api/server"
)

//VersionController responsible for the version resource
type VersionController struct {
	server.PostNotSupported
	server.PutNotSupported
	server.DeleteNotSupported
}

//Get returns the current version of the application
func (versionController *VersionController) Get(values url.Values) (int, interface{}) {
	data := map[string]string{"version": "0.1.0"}
	return 200, data
}
