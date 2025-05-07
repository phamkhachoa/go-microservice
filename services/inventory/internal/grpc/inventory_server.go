package grpc

import (
	"context"
	"go-ecommerce-backend-api/global"
	inventoryGrpc "go-ecommerce-backend-api/proto"
	"google.golang.org/grpc"
	"time"

	"go-ecommerce-backend-api/internal/service"

	"go.uber.org/zap"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// InventoryServer implements the gRPC InventoryService
type InventoryServer struct {
	inventoryGrpc.UnimplementedInventoryServiceServer
	inventoryService service.IInventoryService
}

func NewGrpcInventoryService(grpc *grpc.Server, inventoryService service.IInventoryService) {
	gRPCHandler := &InventoryServer{
		inventoryService: inventoryService,
	}

	// register the OrderServiceServer
	inventoryGrpc.RegisterInventoryServiceServer(grpc, gRPCHandler)
}

// NewInventoryServer creates a new inventory gRPC server
func NewInventoryServer(inventoryService service.IInventoryService) *InventoryServer {
	return &InventoryServer{
		inventoryService: inventoryService,
	}
}

// GetInventoryByProductID implements the gRPC method to retrieve inventory by product ID
func (s *InventoryServer) GetInventoryByProductID(ctx context.Context, req *inventoryGrpc.GetInventoryRequest) (*inventoryGrpc.InventoryResponse, error) {
	global.Logger.Info("Received request for inventory by product ID", zap.Int64("product_id", req.ProductId))

	// Set a timeout for the operation
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	// Get inventory from service
	inventory, err := s.inventoryService.GetByProductId(req.ProductId)
	if err != nil {
		global.Logger.Error("Failed to get inventory", zap.Error(err), zap.Int64("product_id", req.ProductId))
		if err.Error() == "inventory_not_found" {
			return nil, status.Error(codes.NotFound, "Inventory not found for the specified product")
		}
		return nil, status.Error(codes.Internal, "Failed to retrieve inventory information")
	}

	// Format last restock date if available
	var lastRestockDate string
	if inventory.LastRestockDate != nil {
		lastRestockDate = inventory.LastRestockDate.Format(time.RFC3339)
	}

	// Convert to gRPC response
	response := &inventoryGrpc.InventoryResponse{
		Id:                inventory.ID,
		ProductId:         inventory.ProductID,
		Quantity:          int32(inventory.Quantity),
		ReservedQuantity:  int32(inventory.ReservedQuantity),
		AvailableQuantity: int32(inventory.Quantity - inventory.ReservedQuantity),
		IsLowStock:        false,
		LastRestockDate:   lastRestockDate,
	}

	// Set reorder point and quantity if available
	if inventory.ReorderPoint != nil {
		response.ReorderPoint = int32(*inventory.ReorderPoint)
		response.IsLowStock = inventory.Quantity <= *inventory.ReorderPoint
	}

	if inventory.ReorderQuantity != nil {
		response.ReorderQuantity = int32(*inventory.ReorderQuantity)
	}

	global.Logger.Info("Successfully retrieved inventory",
		zap.Int64("product_id", req.ProductId),
		zap.Int32("quantity", response.Quantity),
		zap.Int32("available", response.AvailableQuantity))

	return response, nil
}
