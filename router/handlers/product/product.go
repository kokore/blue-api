package product

import (
	"blue-api/internal/config"
	"blue-api/internal/response"
	"blue-api/internal/service/product"

	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	appConfig      *config.AppConfig
	productService product.Service
}

func Init(appConfig *config.AppConfig, productService product.Service) *ProductHandler {
	return &ProductHandler{
		appConfig:      appConfig,
		productService: productService,
	}
}

func (handler ProductHandler) GetProducts(ginCtx *gin.Context) {

	response.HandlerSuccessResponse(ginCtx, "success")
}
