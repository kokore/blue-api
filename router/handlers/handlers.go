package handlers

import (
	"blue-api/internal/config"
	"blue-api/internal/repository"
	"blue-api/internal/service"
	"blue-api/router/handlers/product"
)

type Handlers struct {
	ProductHandler *product.ProductHandler
}

func Init(appConfig *config.AppConfig, repos repository.Repositories, services service.Services) Handlers {

	return Handlers{
		ProductHandler: product.Init(appConfig, repos),
	}
}