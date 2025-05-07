package repo

import (
	"go-ecommerce-backend-api/internal/model"
)

// ProductRepo defines the interface for product repository operations
type IInventoryRepo interface {
	GetByProductId(id int64) (*model.Inventory, error)
}
