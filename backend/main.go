package main

import (
	"github.com/gin-gonic/gin"
	"github.com/hyperremix/economy-analyzer/backend/api"
)

func main() {
	router := gin.Default()
	api := api.New(router)
	api.RegisterControllers()
	api.Start(3000)
}
