package router

import "go-ecommerce-backend-api/internal/router/product"

type RouterGroup struct {
	Product product.ProductRouterGroup
}

var RouterGroupApp = new(RouterGroup)
