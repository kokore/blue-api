package product

import (
	"blue-api/internal/errorinternal"
	"blue-api/internal/repository/product"

	"github.com/gin-gonic/gin"
)

type CreateProductRequest struct {
	Name         string `json:"name"`
	Price        uint   `json:"price"`
	CurrentStock uint   `json:"currentStock"`
	Image        string `json:"image"`
}

type CreateProductResponseDTO struct {
	Name         string `json:"name"`
	Price        uint   `json:"price"`
	CurrentStock uint   `json:"currentStock"`
	Image        string `json:"image"`
}

type ProductsResponseDTO struct {
	Products []product.Product `json:"products"`
}

func (req CreateProductRequest) Validate(ginCtx *gin.Context) error {
	if req.Name == "" {
		return errorinternal.NewError(errorinternal.ErrorCodeInvalidRequest, "name is required.")
	}
	if req.Price <= 0 {
		return errorinternal.NewError(errorinternal.ErrorCodeInvalidRequest, "price more then 0.")
	}
	if req.CurrentStock <= 0 {
		return errorinternal.NewError(errorinternal.ErrorCodeInvalidRequest, "current stock more then 0.")
	}
	if req.Image == "" {
		return errorinternal.NewError(errorinternal.ErrorCodeInvalidRequest, "image is required.")
	}
	return nil
}
