package controllers

import (
	"net/http"

	"example.com/m/v2/adapters/inbound/dtos"
	"example.com/m/v2/application/ports"
	"github.com/gin-gonic/gin"
)

type ViaCepController struct {
	cep ports.ViaCepClient
}

func NewViaCepClient(apiGroup *gin.RouterGroup, cep ports.ViaCepClient) {
	controller := ViaCepController{cep: cep}
	apiGroup.GET("/viacep/:cep", controller.process())
}

func (controller *ViaCepController) process() gin.HandlerFunc {
	return func(c *gin.Context) {
		cep := c.Param("cep")
		address, err := controller.cep.FindAddressByCep(cep)
		if err != nil {
			c.JSON(http.StatusBadRequest, dtos.Message{
				Message: err.Error(),
				Code:    http.StatusBadRequest,
			})
			return
		}
		if address == nil {
			c.JSON(http.StatusNotFound, "")
			return
		}
		c.JSON(http.StatusOK, address)
	}
}
