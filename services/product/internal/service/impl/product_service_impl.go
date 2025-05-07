package impl

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"go-ecommerce-backend-api/internal/client"
	"time"

	"go-ecommerce-backend-api/global"
	"go-ecommerce-backend-api/internal/common"
	"go-ecommerce-backend-api/internal/model"
	"go-ecommerce-backend-api/internal/repo"
	"go-ecommerce-backend-api/internal/request"
	"go-ecommerce-backend-api/internal/service"
	"go-ecommerce-backend-api/internal/utils"

	"gorm.io/gorm"
)

type productService struct {
	productRepo     repo.ProductRepo
	inventoryClient *client.InventoryClient
}

// NewProductService creates a new product service
func NewProductService(productRepo repo.ProductRepo, inventoryClient *client.InventoryClient) service.IProductService {
	return &productService{
		productRepo:     productRepo,
		inventoryClient: inventoryClient,
	}
}

// CreateProduct creates a new product
func (s *productService) CreateProduct(req *request.CreateProductRequest) (*model.Product, error) {
	product := &model.Product{
		Name:          req.Name,
		Description:   req.Description,
		Price:         req.Price,
		DiscountPrice: req.DiscountPrice,
		Quantity:      req.Quantity,
		CategoryID:    req.CategoryID,
		Thumbnail:     req.Thumbnail,
		Images:        model.JSONMap(req.Images),
		Attributes:    model.JSONMap(req.Attributes),
		Status:        req.Status,
	}

	err := s.productRepo.Create(product)
	if err != nil {
		return nil, err
	}

	return product, nil
}

// GetProductByID retrieves a product by ID
func (s *productService) GetProductByID(id int64) (*model.Product, error) {
	cacheKey := fmt.Sprintf("product:%d", id)
	ctx := context.Background()

	// Try to get the product from Redis cache
	cachedProduct, err := global.Rdb.Get(ctx, cacheKey).Result()
	if err == nil {
		// Product found in cache, unmarshal and return
		var product model.Product
		err = json.Unmarshal([]byte(cachedProduct), &product)
		if err == nil {
			return &product, nil
		}
	}

	// If not in cache or error occurred, fetch from repository
	product, err := s.productRepo.GetByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("not_found")
		}
		return nil, err
	}

	// Get inventory information via gRPC
	inventoryResp, err := s.inventoryClient.GetProductInventory(ctx, id)
	if inventoryResp != nil && err == nil {
		product.Quantity = int(inventoryResp.Quantity)
	}

	// Cache the product in Redis
	productJSON, err := json.Marshal(product)
	if err == nil {
		global.Rdb.Set(ctx, cacheKey, productJSON, 1*time.Hour) // Cache for 1 hour
	}

	return product, nil
}

// UpdateProduct updates a product
func (s *productService) UpdateProduct(id int64, req *request.UpdateProductRequest) (*model.Product, error) {
	product, err := s.productRepo.GetByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("product not found")
		}
		return nil, err
	}

	// Update fields if provided
	if req.Name != "" {
		product.Name = req.Name
	}
	if req.Description != "" {
		product.Description = req.Description
	}
	if req.Price != nil {
		product.Price = *req.Price
	}
	product.DiscountPrice = req.DiscountPrice
	if req.Quantity != nil {
		product.Quantity = *req.Quantity
	}
	if req.CategoryID != nil {
		product.CategoryID = req.CategoryID
	}
	if req.Thumbnail != "" {
		product.Thumbnail = req.Thumbnail
	}
	if req.Images != nil {
		product.Images = model.JSONMap(req.Images)
	}
	if req.Attributes != nil {
		product.Attributes = model.JSONMap(req.Attributes)
	}
	if req.Status != nil {
		product.Status = *req.Status
	}

	err = s.productRepo.Update(product)
	if err != nil {
		return nil, err
	}

	// Invalidate cache
	cacheKey := fmt.Sprintf("product:%d", product.ID)
	global.Rdb.Del(context.Background(), cacheKey)

	return product, nil
}

// DeleteProduct deletes a product
func (s *productService) DeleteProduct(id int64) error {
	_, err := s.productRepo.GetByID(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return errors.New("product not found")
		}
		return err
	}

	return s.productRepo.Delete(id)
}

// SearchProducts searches for products based on filters
func (s *productService) SearchProducts(filter *request.ProductFilter, pagination common.PaginationRequest) ([]model.Product, int64, error) {
	return s.productRepo.Search(filter, pagination)
}

// DeleteMultipleProducts deletes multiple products by their IDs
func (s *productService) DeleteMultipleProducts(ids []int64) error {
	// Validate that we have at least one ID
	if len(ids) == 0 {
		return errors.New("no product IDs provided")
	}

	// Check if all products exist before attempting to delete
	for _, id := range ids {
		_, err := s.productRepo.GetByID(id)
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return fmt.Errorf("product with ID %d not found", id)
			}
			return err
		}
	}

	// Call repository to delete all products in a transaction
	return s.productRepo.DeleteMultiple(ids)
}

// CreateProductWithDetails demonstrates a transactional operation
func (s *productService) CreateProductWithDetails(ctx context.Context, product *model.Product, details []model.ProductDetail) error {
	// Use the transactional decorator
	transactionalFn := utils.Transactional(global.Mdb)(func(ctx context.Context) error {
		// Create the product
		if err := s.productRepo.Create(product); err != nil {
			return err
		}

		// Create product details (assuming you have a repository for product details)
		// This is just an example, you would need to implement the actual repository
		for i := range details {
			details[i].ProductID = product.ID
			// detailRepo.Create(ctx, &details[i])
		}

		return nil
	})

	// Execute the transactional function
	return transactionalFn(ctx)
}
