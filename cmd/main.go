package main

import (
	"example.com/m/v2/adapters/configurations"
	"example.com/m/v2/adapters/inbound/controllers"
	"example.com/m/v2/adapters/inbound/rest"
	"example.com/m/v2/adapters/outbound/repositories"
	"example.com/m/v2/application/ports"
	"example.com/m/v2/application/services"
)

func main() {

	postgresConfig := configurations.NewPostgresConfiguration()
	repository := repositories.NewEmployeeRepositoryImpl(postgresConfig)

	client := rest.NewViaCepClient()
	employeeService := services.NewEmployeeServiceImpl(repository, client)

	go httpServer(client, employeeService)
	flag := make(chan bool)
	<-flag
}

func httpServer(client ports.ViaCepClient, employeeService ports.EmployeeService) {
	ginServer := configurations.NewHttpGinServer()
	apiRoute := ginServer.Group("/api")

	controllers.NewHealthController(ginServer)
	controllers.NewEmployeeController(apiRoute, employeeService)
	controllers.NewViaCepClient(apiRoute, client)

	ginServer.Run(":8080")

}
