//go:build wireinject

package wire

import (
	"go-ecommerce-backend-api/internal/controller"
	repoImpl "go-ecommerce-backend-api/internal/repo/impl"
	serviceImpl "go-ecommerce-backend-api/internal/service/impl"

	"github.com/google/wire"
)

func InitProductRouterHandler() (*controller.ProductController, error) {
	wire.Build(
		repoImpl.NewProductRepo,
		serviceImpl.NewProductService,
		controller.NewProductController)
	return new(controller.ProductController), nil
}
