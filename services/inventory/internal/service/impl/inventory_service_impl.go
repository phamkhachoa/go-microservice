package impl

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"go-ecommerce-backend-api/global"
	"go-ecommerce-backend-api/internal/model"
	"go-ecommerce-backend-api/internal/repo"
	"go-ecommerce-backend-api/internal/service"
	"gorm.io/gorm"
)

type inventoryService struct {
	inventoryRepo repo.IInventoryRepo
}

func NewInventoryService(inventoryRepo repo.IInventoryRepo) service.IInventoryService {
	return &inventoryService{
		inventoryRepo: inventoryRepo,
	}
}

func (s inventoryService) GetByProductId(id int64) (*model.Inventory, error) {
	cacheKey := fmt.Sprintf("inventory:product:%d", id)
	ctx := context.Background()

	// Try to get the product from Redis cache
	cachedInventory, err := global.Rdb.Get(ctx, cacheKey).Result()
	if err == nil {
		// Product found in cache, unmarshal and return
		var inventory model.Inventory
		err = json.Unmarshal([]byte(cachedInventory), &inventory)
		if err == nil {
			return &inventory, nil
		}
	}

	// If not in cache or error occurred, fetch from repository
	inventory, err := s.inventoryRepo.GetByProductId(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("not_found")
		}
		return nil, err
	}

	// Cache the product in Redis
	productJSON, err := json.Marshal(inventory)
	if err == nil {
		global.Rdb.Set(ctx, cacheKey, productJSON, 1*time.Hour) // Cache for 1 hour
	}

	return inventory, nil
}
