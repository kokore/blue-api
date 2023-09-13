package repository

import (
	"blue-api/internal/config"
	"blue-api/internal/database"
	"blue-api/internal/repository/product"
)

type Repositories struct {
	Product product.Repo
}

func Init(appConfig *config.AppConfig, connect database.Connection) Repositories {
	productRepo := product.InitUserRepository(connect)

	return Repositories{
		Product: productRepo,
	}
}
