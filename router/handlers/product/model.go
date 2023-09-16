package product

import (
	"blue-api/internal/errorinternal"
	"blue-api/internal/repository/product"

	"github.com/gin-gonic/gin"
)

type ProductRequest struct {
	Name     string `json:"name"`
	Price    uint   `json:"price"`
	Quantity uint   `json:"quantity"`
	Image    string `json:"image"`
}

type ProductResponseDTO struct {
	Name     string `json:"name"`
	Price    uint   `json:"price"`
	Quantity uint   `json:"quantity"`
	Image    string `json:"image"`
}

type ProductsResponseDTO struct {
	Products []product.Product `json:"products"`
}

func (req ProductRequest) Validate(ginCtx *gin.Context) error {
	if req.Name == "" {
		return errorinternal.NewError(errorinternal.ErrorCodeInvalidRequest, "name is required.")
	}
	if req.Price <= 0 {
		return errorinternal.NewError(errorinternal.ErrorCodeInvalidRequest, "price more then 0.")
	}
	if req.Image == "" {
		return errorinternal.NewError(errorinternal.ErrorCodeInvalidRequest, "image is required.")
	}
	return nil
}
