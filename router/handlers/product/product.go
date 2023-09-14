package product

import (
	"blue-api/internal/config"
	"blue-api/internal/response"
	"blue-api/internal/service/product"
	"net/http"

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

func (handler ProductHandler) CreateProduct(ginCtx *gin.Context) {
	var product ProductRequest
	if err := ginCtx.ShouldBindJSON(&product); err != nil {
		ginCtx.JSON(http.StatusBadRequest, response.Err(response.InvalidRequestJSONString, http.StatusBadRequest, err.Error()))
		return
	}

	err := product.Validate(ginCtx)
	if err != nil {
		ginCtx.JSON(http.StatusBadRequest, response.Err(response.InvalidRequestJSONString, http.StatusBadRequest, err.Error()))
		return
	}

	errService := handler.productService.CreateProductService(ginCtx, product.Name, product.Price, product.CurrentStock, product.Image)
	if errService != nil {
		ginCtx.JSON(http.StatusBadRequest, response.Err(response.UnableInquiryProduct, http.StatusBadRequest, errService.Error()))
		return
	}

	response.HandlerSuccessResponse(ginCtx, ToCreatedProductResponseDTO(&product))
}

func (handler ProductHandler) GetProducts(ginCtx *gin.Context) {
	products, err := handler.productService.GetProductsService(ginCtx)
	if err != nil {
		ginCtx.JSON(http.StatusBadRequest, response.Err(response.UnableGetProduct, http.StatusBadRequest, err.Error()))
		return
	}
	response.HandlerSuccessResponse(ginCtx, ToProductsResponseDTO(products))
}

func (handler ProductHandler) UpdateProduct(ginCtx *gin.Context) {
	productId := ginCtx.Param("product_id")

	var product ProductRequest
	if err := ginCtx.ShouldBindJSON(&product); err != nil {
		ginCtx.JSON(http.StatusBadRequest, response.Err(response.InvalidRequestJSONString, http.StatusBadRequest, err.Error()))
		return
	}

	err := handler.productService.UpdateProductServie(ginCtx, productId, product.Name, product.Price, product.CurrentStock, product.Image)
	if err != nil {
		ginCtx.JSON(http.StatusBadRequest, response.Err(response.UnableInquiryProduct, http.StatusBadRequest, err.Error()))
		return
	}

	response.HandlerSuccessResponse(ginCtx, ToUpdateProductResponseDTO(&product))
}
