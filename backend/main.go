package main

import (
	"github.com/gin-gonic/gin"
	"github.com/hyperremix/economy-analyzer/backend/api/server"
)

func main() {
	router := gin.Default()
	api := server.NewAPI(router)
	api.RegisterControllers()
	api.Start(3000)
}
