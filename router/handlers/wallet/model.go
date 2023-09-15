package wallet

import (
	"blue-api/internal/errorinternal"

	"github.com/gin-gonic/gin"
)

type WalletRequest struct {
	Coins     []int `json:"coins"`
	Banknotes []int `json:"banknotes"`
}

func (req WalletRequest) Validate(ginCtx *gin.Context) error {
	if len(req.Coins) > 3 {
		return errorinternal.NewError(errorinternal.ErrorCodeInvalidRequest, "invalid coins.")
	}
	if len(req.Banknotes) > 5 {
		return errorinternal.NewError(errorinternal.ErrorCodeInvalidRequest, "invalid banknotes.")
	}
	return nil
}
