package repository

import (
	"blue-api/internal/config"
	"blue-api/internal/database"
	"blue-api/internal/repository/product"
	"blue-api/internal/repository/wallet"
)

type Repositories struct {
	Product product.Repo
	Wallet  wallet.Repo
}

func Init(appConfig *config.AppConfig, connect database.Connection) Repositories {
	productRepo := product.InitProductRepository(connect)
	walletRepo := wallet.InitWalletRepository(connect)

	return Repositories{
		Product: productRepo,
		Wallet:  walletRepo,
	}
}
