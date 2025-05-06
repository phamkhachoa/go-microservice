package product

import (
	"go-ecommerce-backend-api/internal/wire"

	"github.com/gin-gonic/gin"
)

type ProductRouter struct {
}

// InitRouter initializes the product routes
func (pr *ProductRouter) InitRouter(Router *gin.RouterGroup) {

	productController, _ := wire.InitProductRouterHandler()

	// Public routes
	productRouterPublic := Router.Group("/products")
	{
		productRouterPublic.GET("/search", productController.SearchProducts)
		productRouterPublic.GET("/:id", productController.GetProduct)
	}

	// Private routes (require authentication)
	productRouterPrivate := Router.Group("/products")
	// Add middleware for authentication if needed
	// productRouterPrivate.Use(middleware.AuthMiddleware())
	{
		productRouterPrivate.POST("", productController.CreateProduct)
		productRouterPrivate.PUT("/:id", productController.UpdateProduct)
		productRouterPrivate.DELETE("/:id", productController.DeleteProduct)
		productRouterPrivate.POST("/batch-delete", productController.DeleteMultipleProducts) // New endpoint for batch deletion
	}
}
