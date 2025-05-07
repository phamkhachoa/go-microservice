package impl

import (
	"go-ecommerce-backend-api/global"
	"go-ecommerce-backend-api/internal/model"
	"go-ecommerce-backend-api/internal/repo"
)

type inventoryRepo struct {
	//db *gorm.DB
}

// NewProductRepo creates a new product repository
func NewInventoryRepo() repo.IInventoryRepo {
	return &inventoryRepo{}
}

func (i inventoryRepo) GetByProductId(id int64) (*model.Inventory, error) {
	var inventory model.Inventory
	err := global.Mdb.Where("product_id = ?", id).First(&inventory).Error
	if err != nil {
		return nil, err
	}
	return &inventory, nil
}
