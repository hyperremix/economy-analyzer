package version

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const path = "/info/version"

func RegisterVersionController(router *gin.Engine, routePrefix string) {
	router.GET(routePrefix+path, get())
}

func get() gin.HandlerFunc {
	return func(c *gin.Context) {
		data := map[string]string{"version": "0.1.0"}
		c.JSON(http.StatusOK, data)
	}
}
