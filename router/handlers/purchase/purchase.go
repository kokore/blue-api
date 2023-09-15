package purchase

import (
	"blue-api/internal/config"
	"blue-api/internal/response"
	"blue-api/internal/service/purchase"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PurchaseHandler struct {
	appConfig       *config.AppConfig
	purchaseService purchase.Service
}

func Init(appConfig *config.AppConfig, purchaseService purchase.Service) *PurchaseHandler {
	return &PurchaseHandler{
		appConfig:       appConfig,
		purchaseService: purchaseService,
	}
}

func (handler *PurchaseHandler) Purchase(ginCtx *gin.Context) {
	var purchase PurchaseRequest
	if err := ginCtx.ShouldBindJSON(&purchase); err != nil {
		ginCtx.JSON(http.StatusBadRequest, response.Err(response.InvalidRequestJSONString, http.StatusBadRequest, err.Error()))
		return
	}

	err := purchase.Validate(ginCtx)
	if err != nil {
		ginCtx.JSON(http.StatusBadRequest, response.Err(response.InvalidRequestJSONString, http.StatusBadRequest, err.Error()))
		return
	}

	errService := handler.purchaseService.PurchaseProcessService(ginCtx, purchase.ProductId, purchase.Quantity)
	if errService != nil {
		ginCtx.JSON(http.StatusBadRequest, response.Err(response.UnableInquiryPurchase, http.StatusBadRequest, errService.Error()))
		return
	}
	response.HandlerSuccessResponse(ginCtx, "success")
}
