package wallet

import (
	"blue-api/internal/config"
	"blue-api/internal/response"
	"blue-api/internal/service/wallet"
	"net/http"

	"github.com/gin-gonic/gin"
)

type WalletHandler struct {
	appConfig     *config.AppConfig
	walletService wallet.Service
}

func Init(appConfig *config.AppConfig, walletService wallet.Service) *WalletHandler {
	return &WalletHandler{
		appConfig:     appConfig,
		walletService: walletService,
	}
}

func (handler *WalletHandler) CreateWallet(ginCtx *gin.Context) {
	var wallet WalletRequest
	if err := ginCtx.ShouldBindJSON(&wallet); err != nil {
		ginCtx.JSON(http.StatusBadRequest, response.Err(response.InvalidRequestJSONString, http.StatusBadRequest, err.Error()))
		return
	}

	err := wallet.Validate(ginCtx)
	if err != nil {
		ginCtx.JSON(http.StatusBadRequest, response.Err(response.InvalidRequestJSONString, http.StatusBadRequest, err.Error()))
		return
	}

	errService := handler.walletService.CreateWalletService(ginCtx, wallet.Coins, wallet.Banknotes)
	if errService != nil {
		ginCtx.JSON(http.StatusBadRequest, response.Err(response.UnableInquiryProduct, http.StatusBadRequest, errService.Error()))
		return
	}

	response.HandlerSuccessResponse(ginCtx, "success")
}
