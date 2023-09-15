package purchase

import (
	"blue-api/internal/config"
	"blue-api/internal/repository"
	productRepo "blue-api/internal/repository/product"
	walletRepo "blue-api/internal/repository/wallet"
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type serviceImpl struct {
	appConfig         *config.AppConfig
	productRepository productRepo.Repo
	walletRepository  walletRepo.Repo
}

type Service interface {
	PurchaseProcessService(ctx context.Context, productId string, quantity int) error
}

func InitPurchaseService(appConfig *config.AppConfig, repos repository.Repositories) Service {
	return &serviceImpl{
		appConfig:         appConfig,
		productRepository: repos.Product,
		walletRepository:  repos.Wallet,
	}
}

func distributeAmount(total int) (map[int]int, map[int]int) {
	coinDenominations := []int{10, 5, 1}
	banknoteDenominations := []int{1000, 500, 100, 50, 20}

	dataCoins := make(map[int]int)
	dataBanknotes := make(map[int]int)

	for _, banknote := range banknoteDenominations {
		dataBanknotes[banknote] = total / banknote
		total %= banknote
	}

	for _, coin := range coinDenominations {
		dataCoins[coin] = total / coin
		total %= coin
	}

	dataBanknotes[banknoteDenominations[len(banknoteDenominations)-1]] += total / banknoteDenominations[len(banknoteDenominations)-1]

	return dataCoins, dataBanknotes
}

func (s serviceImpl) PurchaseProcessService(ctx context.Context, productId string, quantity int) error {
	objectID, err := primitive.ObjectIDFromHex(productId)
	if err != nil {
		return err
	}

	product, err := s.productRepository.FindById(ctx, objectID)
	if err != nil {
		return err
	}

	if product.Quantity == 0 {
		return errors.New("out of stock")
	}

	if product.Quantity-quantity < 0 {
		return errors.New("out of stock")
	}

	wallet, err := s.walletRepository.FindOne(ctx)
	if err != nil {
		return err
	}

	totalPrice := quantity * product.Price

	if wallet.Total-totalPrice < 0 {
		return errors.New("not enough money")
	}

	totalAmount := wallet.Total - totalPrice
	totalQuantity := product.Quantity - quantity

	dataCoins, dataBanknotes := distributeAmount(totalAmount)

	_, errProduct := s.productRepository.FindAndUpdate(ctx, product.ID, productRepo.NewUpdate().SetQuantity(uint(totalQuantity)))
	if errProduct != nil {
		return errProduct
	}
	_, errWallet := s.walletRepository.UpdateOne(ctx, wallet.ID, walletRepo.NewUpdate().SetCoins(dataCoins).SetBanknotes(dataBanknotes).SetTotal(totalAmount))
	if errWallet != nil {
		return errWallet
	}
	return nil
}
