package inventory

import (
	"go-ecommerce-backend-api/internal/wire"

	"github.com/gin-gonic/gin"
)

type InventoryRouter struct {
}

// InitRouter initializes the product routes
func (pr *InventoryRouter) InitRouter(Router *gin.RouterGroup) {

	inventoryController, _ := wire.InitInventoryRouterHandler()

	// Public routes
	productRouterPublic := Router.Group("/inventories")
	{
		productRouterPublic.GET("/products/:id", inventoryController.GetByProductId)
	}
}
