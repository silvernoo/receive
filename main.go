package main

import (
	"flag"
	"github.com/gin-gonic/gin"
	"os"
	"receive/handler"
	"receive/initRouter"
)

var (
	h string
)

func init() {
	flag.StringVar(&h, "p", "upload", "Set path")
}

func main() {
	if os.Getenv("DEBUG") == "debug" {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}
	flag.Parse()
	handler.Path = h
	router := initRouter.SetupRouter()
	_ = router.Run(":80")
}
