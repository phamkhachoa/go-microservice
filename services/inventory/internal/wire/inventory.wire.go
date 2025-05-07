//go:build wireinject

package wire

import (
	"go-ecommerce-backend-api/internal/controller"
	repoImpl "go-ecommerce-backend-api/internal/repo/impl"
	serviceImpl "go-ecommerce-backend-api/internal/service/impl"

	"github.com/google/wire"
)

func InitInventoryRouterHandler() (*controller.InventoryController, error) {
	wire.Build(
		repoImpl.NewInventoryRepo,
		serviceImpl.NewInventoryService,
		controller.NewInventoryController)
	return new(controller.InventoryController), nil
}
