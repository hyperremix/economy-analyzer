package main

import (
	"github.com/hyperremix/economy-analyzer/api"
	"github.com/hyperremix/economy-analyzer/version"
)

func main() {

	versionController := new(version.VersionController)

	var api = new(api.API)
	api.AddController(versionController, "/info/version")
	api.Start(3000)
}
