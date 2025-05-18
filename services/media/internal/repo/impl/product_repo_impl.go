package impl

import (
	"fmt"
	"go-ecommerce-backend-api/global"
	"go-ecommerce-backend-api/internal/common"
	"go-ecommerce-backend-api/internal/model"
	"go-ecommerce-backend-api/internal/repo"
	"go-ecommerce-backend-api/internal/request"
	"strings"
)

type productRepo struct {
	//db *gorm.DB
}

// NewProductRepo creates a new product repository
func NewProductRepo() repo.ProductRepo {
	return &productRepo{}
}

// Create creates a new product
func (r *productRepo) Create(product *model.Product) error {
	return global.Mdb.Create(product).Error
}

// GetByID retrieves a product by ID
func (r *productRepo) GetByID(id int64) (*model.Product, error) {
	var product model.Product
	err := global.Mdb.Where("id = ? AND status != -1", id).First(&product).Error
	if err != nil {
		return nil, err
	}
	return &product, nil
}

// Update updates a product
func (r *productRepo) Update(product *model.Product) error {
	return global.Mdb.Save(product).Error
}

// Delete soft deletes a product
func (r *productRepo) Delete(id int64) error {
	return global.Mdb.Model(&model.Product{}).Where("id = ?", id).Update("status", -1).Error
}

// Search searches for products based on filters
func (r *productRepo) Search(filter *request.ProductFilter, pagination common.PaginationRequest) ([]model.Product, int64, error) {
	var products []model.Product

	query := global.Mdb.Model(&model.Product{}).Where("status != -1")

	// Apply filters
	if filter.Name != "" {
		query = query.Where("name LIKE ?", "%"+filter.Name+"%")
	}

	if filter.CategoryID != nil {
		query = query.Where("category_id = ?", *filter.CategoryID)
	}

	if filter.MinPrice != nil {
		query = query.Where("price >= ?", *filter.MinPrice)
	}

	if filter.MaxPrice != nil {
		query = query.Where("price <= ?", *filter.MaxPrice)
	}

	if filter.Status != nil {
		query = query.Where("status = ?", *filter.Status)
	}

	// Count total records for pagination
	var totalItems int64
	query.Count(&totalItems)

	// Apply sorting
	if pagination.Sort != "" {
		parts := strings.Split(pagination.Sort, ":")
		if len(parts) == 2 {
			field := parts[0]
			direction := parts[1]

			// Validate direction
			if direction != "asc" && direction != "desc" {
				direction = "asc"
			}

			query = query.Order(fmt.Sprintf("%s %s", field, direction))
		}
	}

	// Apply pagination
	offset := (pagination.Page - 1) * pagination.Limit
	query = query.Offset(int(offset)).Limit(int(pagination.Limit))

	err := query.Find(&products).Error
	return products, totalItems, err
}

// DeleteMultiple soft deletes multiple products by their IDs using a transaction
func (r *productRepo) DeleteMultiple(ids []int64) error {
	// Begin a transaction
	tx := global.Mdb.Begin()
	if tx.Error != nil {
		return tx.Error
	}

	// Defer a function to handle rollback/commit based on error
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// Perform the delete operation within the transaction
	if err := tx.Model(&model.Product{}).Where("id IN ?", ids).Update("status", -1).Error; err != nil {
		tx.Rollback()
		return err
	}

	// Commit the transaction
	return tx.Commit().Error
}
