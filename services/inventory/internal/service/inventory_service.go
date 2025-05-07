package service

import (
	"go-ecommerce-backend-api/internal/model"
)

// ProductService defines the interface for product service operations
type IInventoryService interface {
	GetByProductId(id int64) (*model.Inventory, error)
}
