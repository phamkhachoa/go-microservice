package repo

import (
	"go-ecommerce-backend-api/internal/common"
	"go-ecommerce-backend-api/internal/model"
	"go-ecommerce-backend-api/internal/request"
)

// ProductRepo defines the interface for product repository operations
type ProductRepo interface {
	Create(product *model.Product) error
	GetByID(id int64) (*model.Product, error)
	Update(product *model.Product) error
	Delete(id int64) error
	DeleteMultiple(ids []int64) error
	Search(filter *request.ProductFilter, pagination common.PaginationRequest) ([]model.Product, int64, error)
}
