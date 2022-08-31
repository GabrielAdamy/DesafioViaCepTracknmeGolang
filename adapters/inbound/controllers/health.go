package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type HealthController struct {
}

func NewHealthController(ginServer *gin.Engine) {
	controller := HealthController{}
	status := ginServer.Group("/status")
	status.GET("/health", controller.health())
}

func (controller *HealthController) health() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": "UP",
		})
	}
}
