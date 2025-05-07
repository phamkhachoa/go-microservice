package router

import "go-ecommerce-backend-api/internal/router/inventory"

type RouterGroup struct {
	Inventory inventory.InventoryRouterGroup
}

var RouterGroupApp = new(RouterGroup)
