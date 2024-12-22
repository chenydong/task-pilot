package api

import (
	"github.com/gin-gonic/gin"
	"task_pilot/config"
)

func ResisterRouter(ctx *config.ServiceContext) *gin.Engine {

	router := gin.Default()
	router.Use(gin.Recovery(), gin.Logger())
	v1 := router.Group("/v1")
	{
		v1.GET("/ping", func(context *gin.Context) {
			context.JSON(200, "PING SUCCESS")
		})
		v1.GET("/create")
	}

	return router
}
