package purchase

import (
	"blue-api/internal/repository/product"
	"blue-api/internal/repository/wallet"
)

func ToPurchaseResponseDTO(products []product.Product, wallet *wallet.Wallet) *PurchaseResponseDTO {
	return &PurchaseResponseDTO{
		Products: products,
		Wallet:   wallet,
	}
}
