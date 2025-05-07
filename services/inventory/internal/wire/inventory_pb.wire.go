//go:build wireinject

package wire

import (
	"go-ecommerce-backend-api/internal/grpc"
	repoImpl "go-ecommerce-backend-api/internal/repo/impl"
	"go-ecommerce-backend-api/internal/service"
	serviceImpl "go-ecommerce-backend-api/internal/service/impl"

	"github.com/google/wire"
)

// InitInventoryGrpcServer initializes the gRPC server
func InitInventoryServer() (*grpc.InventoryServer, error) {
	wire.Build(
		repoImpl.NewInventoryRepo,
		serviceImpl.NewInventoryService,
		provideInventoryServer,
	)
	return new(grpc.InventoryServer), nil
}

// provideInventoryServer creates a new inventory gRPC server
func provideInventoryServer(inventoryService service.IInventoryService) *grpc.InventoryServer {
	return grpc.NewInventoryServer(inventoryService)
}
