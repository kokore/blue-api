package router

import (
	"blue-api/internal/config"
	"blue-api/router/handlers"
	"blue-api/router/middleware"
	"fmt"

	"github.com/gin-gonic/gin"
)

func Init(appConfig *config.AppConfig, handlers handlers.Handlers) {
	port := fmt.Sprintf(":%s", appConfig.APIConfig.Port)

	engine := gin.Default()

	engine.Use(middleware.CORS(appConfig))

	baseGroup := engine.Group("/api")
	version1 := baseGroup.Group("/v1")

	productRouters := version1.Group("/product")
	registerProductRoutes(productRouters, handlers)

	vendingMachineRouters := version1.Group("/vm")
	registervendingMachineRoutersRoutes(vendingMachineRouters, handlers)

	engine.Run(port)
}

func registerProductRoutes(group *gin.RouterGroup, handlers handlers.Handlers) {
	group.GET("", handlers.ProductHandler.GetProducts)
	group.POST("", handlers.ProductHandler.CreateProduct)
	group.PATCH("/:product_id", handlers.ProductHandler.UpdateProduct)
	group.DELETE("/:product_id", handlers.ProductHandler.DeleteProduct)
}

func registervendingMachineRoutersRoutes(group *gin.RouterGroup, handlers handlers.Handlers) {
	group.GET("")
}
