package response

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandlerSuccessResponse(c *gin.Context, response any) {
	c.JSON(http.StatusOK, response)
}
