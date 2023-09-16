package purchase

import (
	"blue-api/internal/errorinternal"
	"blue-api/internal/repository/product"
	"blue-api/internal/repository/wallet"

	"github.com/gin-gonic/gin"
)

type PurchaseRequest struct {
	ProductId string `json:"productId"`
	Quantity  int    `json:"quantity"`
}

type PurchaseResponseDTO struct {
	Products []product.Product `json:"products"`
	Wallet   *wallet.Wallet    `json:"wallet"`
}

func (req PurchaseRequest) Validate(ginCtx *gin.Context) error {
	if req.ProductId == "" {
		return errorinternal.NewError(errorinternal.ErrorCodeInvalidRequest, "product id is required.")
	}
	if req.Quantity <= 0 {
		return errorinternal.NewError(errorinternal.ErrorCodeInvalidRequest, "quantity more then 0.")
	}
	return nil
}
