package main

import (
	"github.com/hyperremix/economy-analyzer/backend/api/monthly"
	"github.com/hyperremix/economy-analyzer/backend/api/server"
	"github.com/hyperremix/economy-analyzer/backend/api/version"
)

func main() {
	var api = new(server.API)
	api.AddController(new(version.VersionController), "/info/version")
	api.AddController(new(monthly.MonthlyController), "/monthlies")
	api.Start(3000)
}
