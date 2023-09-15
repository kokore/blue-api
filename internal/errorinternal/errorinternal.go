package errorinternal

import "fmt"

const (
	ErrorCodeUnauthorized       = "UNAUTHORIZED"
	ErrorCodeInvalidRequest     = "INVALID_REQUEST"
	ErrorCodeNotFound           = "NOT_FOUND"
	ErrorCodeInternal           = "INTERNAL_SERVER_ERROR"
	ErrorCodeDoNotMapTranslator = "ERROR_CODE_DO_NOT_MAP_TO_TRANSLATOR"

	// business code error
	ErrorCodeProductCantInsert = "PRODUCT_CANT_INSERT"
	ErrorCodeProductNotFound   = "PRODUCT_NOT_FOUND"

	ErrorCodeWalletCantInsert = "WALLE_CANT_INSERT"
	ErrorCodeWalletNotFound   = "WALLET_NOT_FOUND"
)

type Error struct {
	Code    string
	Message string
}

func (err *Error) Error() string {
	return fmt.Sprintf("%s: %s", err.Code, err.Message)
}

func NewError(code, message string) error {
	return &Error{
		Code:    code,
		Message: message,
	}
}
