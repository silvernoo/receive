package initRouter

import (
	"github.com/gin-gonic/gin"
	"receive/handler"
)

func SetupRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger(), gin.Recovery())
	index := router.Group("/")
	{
		index.PUT("/receive", handler.Receive)
	}
	return router
}
