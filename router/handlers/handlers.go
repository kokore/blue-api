package handlers

import (
	"blue-api/internal/config"
	"blue-api/internal/repository"
	"blue-api/internal/service"
	"blue-api/router/handlers/product"
	"blue-api/router/handlers/wallet"
)

type Handlers struct {
	ProductHandler *product.ProductHandler
	WalletHandler  *wallet.WalletHandler
}

func Init(appConfig *config.AppConfig, repos repository.Repositories, services service.Services) Handlers {

	return Handlers{
		ProductHandler: product.Init(appConfig, services.Product),
		WalletHandler:  wallet.Init(appConfig, services.Wallet),
	}
}
