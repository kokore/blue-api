package service

import (
	"blue-api/internal/config"
	"blue-api/internal/repository"
	productSerive "blue-api/internal/service/product"
)

type Services struct {
	Product productSerive.Service
}

func Init(appConfig *config.AppConfig, repos repository.Repositories) Services {
	productSerive := productSerive.InitProductService(appConfig, repos)
	return Services{
		Product: productSerive,
	}
}
