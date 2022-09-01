package configurations

import (
	"os"

	"github.com/gin-gonic/gin"
)

var ACCESS_SECRET string = os.Getenv("ACCESS_SECRET")

func NewHttpGinServer() *gin.Engine {
	ginServer := gin.Default()
	ginServer.Use(gin.Recovery())
	ginServer.Use(gin.Logger())

	return ginServer
}
