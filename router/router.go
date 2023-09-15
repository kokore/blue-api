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

	walletRouters := version1.Group("/wallet")
	registerWalletRoutersRoutes(walletRouters, handlers)

	purchaseRouters := version1.Group("/purchase")
	registerPurchaseRoutes(purchaseRouters, handlers)

	engine.Run(port)
}

func registerProductRoutes(group *gin.RouterGroup, handlers handlers.Handlers) {
	group.GET("", handlers.ProductHandler.GetProducts)
	group.POST("", handlers.ProductHandler.CreateProduct)
	group.PATCH("/:product_id", handlers.ProductHandler.UpdateProduct)
	group.DELETE("/:product_id", handlers.ProductHandler.DeleteProduct)
}

func registerWalletRoutersRoutes(group *gin.RouterGroup, handlers handlers.Handlers) {
	group.GET("", handlers.WalletHandler.GetWallet)
	group.POST("", handlers.WalletHandler.CreateWallet)
	group.PATCH("/:wallet_id", handlers.WalletHandler.UpdateWallet)
}

func registerPurchaseRoutes(group *gin.RouterGroup, handlers handlers.Handlers) {
	group.POST("", handlers.PurchaseHandler.Purchase)
}
