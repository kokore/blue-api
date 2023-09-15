package service

import (
	"blue-api/internal/config"
	"blue-api/internal/repository"
	productSerive "blue-api/internal/service/product"
	purchaseService "blue-api/internal/service/purchase"
	walletService "blue-api/internal/service/wallet"
)

type Services struct {
	Product  productSerive.Service
	Wallet   walletService.Service
	Purchase purchaseService.Service
}

func Init(appConfig *config.AppConfig, repos repository.Repositories) Services {
	productSerive := productSerive.InitProductService(appConfig, repos)
	walletService := walletService.InitWalletService(appConfig, repos)
	purchaseService := purchaseService.InitPurchaseService(appConfig, repos)
	return Services{
		Product:  productSerive,
		Wallet:   walletService,
		Purchase: purchaseService,
	}
}
