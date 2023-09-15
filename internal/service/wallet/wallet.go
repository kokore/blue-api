package wallet

import (
	"blue-api/internal/config"
	"blue-api/internal/repository"
	walletRepo "blue-api/internal/repository/wallet"
	"context"
)

type serviceImpl struct {
	appConfig        *config.AppConfig
	walletRepository walletRepo.Repo
}

type Service interface {
	CreateWalletService(ctx context.Context, coins []int, backnotes []int) error
}

func InitWalletService(appConfig *config.AppConfig, repos repository.Repositories) Service {
	return &serviceImpl{
		appConfig:        appConfig,
		walletRepository: repos.Wallet,
	}
}

func calculateTotal(dataMap map[int]int) int {
	total := 0
	for value, count := range dataMap {
		total += value * count
	}
	return total
}

func (s serviceImpl) CreateWalletService(ctx context.Context, coins []int, backnotes []int) error {

	dataCoins := map[int]int{
		1:  coins[0],
		5:  coins[1],
		10: coins[2],
	}

	dataBacknotes := map[int]int{
		20:   backnotes[0],
		50:   backnotes[1],
		100:  backnotes[2],
		500:  backnotes[3],
		1000: backnotes[4],
	}

	total := calculateTotal(dataCoins) + calculateTotal(dataBacknotes)

	err := s.walletRepository.InsertOne(ctx, walletRepo.NewFilter().SetCoins(dataCoins).SetBanknotes(dataBacknotes).SetTotal(total))
	if err != nil {
		return err
	}
	return nil
}
