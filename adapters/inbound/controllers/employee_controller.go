package controllers

import (
	"io/ioutil"
	"net/http"
	"strconv"

	"example.com/m/v2/adapters/inbound/dtos"
	"example.com/m/v2/application/domain"
	"example.com/m/v2/application/ports"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type EmployeeController struct {
	service ports.EmployeeService
}

func NewEmployeeController(apiGroup *gin.RouterGroup, service ports.EmployeeService) {
	controller := EmployeeController{service}
	apiGroup.POST("/employee", controller.create())
	apiGroup.GET("/employee/:id", controller.findOne())
	apiGroup.GET("/employee", controller.findAll())
	apiGroup.GET("/employee/cep/:cep", controller.findByCep())
	apiGroup.PUT("/employee", controller.update())
	apiGroup.DELETE("/employee/:id", controller.delete())
	apiGroup.PATCH("/employee/:id", controller.UpdatePatch())

}

func (controller *EmployeeController) create() gin.HandlerFunc {
	return func(c *gin.Context) {
		var employeeDto dtos.EmployeeDTO

		if err := c.ShouldBindJSON(&employeeDto); err != nil {
			c.JSON(http.StatusBadRequest, dtos.Message{
				Message: err.Error(),
				Code:    http.StatusBadRequest,
			})
			return
		}
		employee := employeeDto.ParseToDomain()
		if err := controller.service.Save(employee); err != nil {
			c.JSON(http.StatusBadRequest, dtos.Message{
				Message: err.Error(),
				Code:    http.StatusBadRequest,
			})
			return
		}
		response := dtos.NewContractDTO(employee)
		c.JSON(http.StatusCreated, response)
	}
}

func (controller *EmployeeController) findOne() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, dtos.Message{
				Message: err.Error(),
				Code:    http.StatusBadRequest,
			})
			return
		}
		domain, err := controller.service.FindOne(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, dtos.Message{
				Message: err.Error(),
				Code:    http.StatusBadRequest,
			})
			return
		}
		if domain == nil {
			c.JSON(http.StatusNotFound, "")
			return
		}
		c.JSON(http.StatusOK, domain)
	}
}

func (controller *EmployeeController) findAll() gin.HandlerFunc {
	return func(c *gin.Context) {
		var parameters = make(map[string]string)
		if err := c.ShouldBindWith(&parameters, binding.Query); err != nil {
			c.JSON(http.StatusBadRequest, dtos.Message{
				Message: err.Error(),
				Code:    http.StatusBadRequest,
			})
			return
		}
		rows, limit := controller.getPageRequest(parameters)
		domains, count, err := controller.service.FindAll(rows, limit, parameters)
		if err != nil {
			c.JSON(http.StatusBadRequest, dtos.Message{
				Message: err.Error(),
				Code:    http.StatusBadRequest,
			})
			return
		}
		pageResources := controller.createPage(domains, count, rows, limit)
		c.JSON(http.StatusOK, pageResources)
	}
}

func (controller *EmployeeController) findByCep() gin.HandlerFunc {
	return func(c *gin.Context) {
		cep := c.Param("cep")
		var parameters = make(map[string]string)
		parameters["cep"] = cep
		if err := c.ShouldBindWith(&parameters, binding.Query); err != nil {
			c.JSON(http.StatusBadRequest, dtos.Message{
				Message: err.Error(),
				Code:    http.StatusBadRequest,
			})
			return
		}
		rows, limit := controller.getPageRequest(parameters)
		ceps, count, err := controller.service.FindByCep(cep, rows, limit)
		if err != nil {
			c.JSON(http.StatusBadRequest, dtos.Message{
				Message: err.Error(),
				Code:    http.StatusBadRequest,
			})
			return
		}
		pageResources := controller.createPage(ceps, count, rows, limit)
		c.JSON(http.StatusOK, pageResources)
	}
}

func (controller *EmployeeController) UpdatePatch() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, dtos.Message{
				Message: err.Error(),
				Code:    http.StatusBadRequest,
			})
			return
		}
		json, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			c.JSON(http.StatusBadRequest, dtos.Message{
				Message: err.Error(),
				Code:    http.StatusBadRequest,
			})
			return
		}

		employee := controller.service.UpdatePatch(id, json)
		if err != nil {
			c.JSON(http.StatusBadRequest, dtos.Message{
				Message: err.Error(),
				Code:    http.StatusBadRequest,
			})
			return
		}

		c.JSON(http.StatusOK, &employee)
	}
}

func (controller *EmployeeController) update() gin.HandlerFunc {
	return func(c *gin.Context) {
		var employeeDto dtos.EmployeeDTO

		if err := c.ShouldBindJSON(&employeeDto); err != nil {
			c.JSON(http.StatusBadRequest, dtos.Message{
				Message: err.Error(),
				Code:    http.StatusBadRequest,
			})
			return
		}
		employee := employeeDto.ParseToDomain()
		if err := controller.service.Save(employee); err != nil {
			c.JSON(http.StatusBadRequest, dtos.Message{
				Message: err.Error(),
				Code:    http.StatusBadRequest,
			})
			return
		}
		response := dtos.NewContractDTO(employee)
		c.JSON(http.StatusCreated, response)
	}
}

func (controller *EmployeeController) delete() gin.HandlerFunc {
	return func(c *gin.Context) {
		id, err := strconv.ParseInt(c.Param("id"), 10, 64)
		if err != nil {
			c.JSON(http.StatusBadRequest, dtos.Message{
				Message: err.Error(),
				Code:    http.StatusBadRequest,
			})
			return
		}
		d := controller.service.Delete(id)
		if err != nil {
			c.JSON(http.StatusBadRequest, dtos.Message{
				Message: err.Error(),
				Code:    http.StatusBadRequest,
			})
			return
		}
		c.JSON(http.StatusOK, d)
	}
}

func (controller *EmployeeController) createPage(domains []*domain.Employee, count int64, page int64, limit int64) *dtos.EmployeePage {
	resources := make([]*dtos.EmployeeDTO, len(domains))
	for i, v := range domains {
		resources[i] = dtos.NewContractDTO(v)
	}

	pageContracts := dtos.NewContractPage(resources, count, page, limit)
	return pageContracts
}

func (controller *EmployeeController) getPageRequest(parameters map[string]string) (int64, int64) {
	var page int64 = 0
	var limit int64 = 10

	if parameters["page"] != "" {
		page, _ = strconv.ParseInt(parameters["page"], 10, 64)
	}

	if parameters["limit"] != "" {
		limit, _ = strconv.ParseInt(parameters["limit"], 10, 64)
	}
	return page, limit
}
