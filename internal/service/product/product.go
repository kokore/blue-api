package product

import (
	"blue-api/internal/config"
	"blue-api/internal/repository"
	productRepo "blue-api/internal/repository/product"
)

type serviceImpl struct {
	appConfig         *config.AppConfig
	productRepository productRepo.Repo
}

type Service interface {
}

func InitProductService(appConfig *config.AppConfig, repos repository.Repositories) Service {
	return &serviceImpl{
		appConfig:         appConfig,
		productRepository: repos.Product,
	}
}
