package main

import (
	"blue-api/internal/config"
	"blue-api/internal/database"
	"blue-api/internal/repository"
	"blue-api/internal/service"
	"blue-api/router"
	"blue-api/router/handlers"
	"fmt"
)

func main() {
	appConfig := config.LoadConfig()

	databaseConnect, err := database.NewConnection(&appConfig.DatabaseConfig)
	if err != nil {
		fmt.Println("Error connect database.")
		panic(err)
	}

	repositories := repository.Init(appConfig, databaseConnect)
	services := service.Init(appConfig, repositories)
	handlers := handlers.Init(appConfig, repositories, services)

	router.Init(appConfig, handlers)
}
