package wallet

import (
	"blue-api/internal/config"
	"blue-api/internal/repository"
	walletRepo "blue-api/internal/repository/wallet"
)

type serviceImpl struct {
	appConfig        *config.AppConfig
	walletRepository walletRepo.Repo
}

type Service interface {
}

func InitWalletService(appConfig *config.AppConfig, repos repository.Repositories) Service {
	return &serviceImpl{
		appConfig:        appConfig,
		walletRepository: repos.Wallet,
	}
}
