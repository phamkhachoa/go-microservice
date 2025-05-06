package service

import (
	"context"
	"go-ecommerce-backend-api/internal/common"
	"go-ecommerce-backend-api/internal/model"
	"go-ecommerce-backend-api/internal/request"
)

// ProductService defines the interface for product service operations
type IProductService interface {
	CreateProduct(req *request.CreateProductRequest) (*model.Product, error)
	GetProductByID(id int64) (*model.Product, error)
	UpdateProduct(id int64, req *request.UpdateProductRequest) (*model.Product, error)
	DeleteProduct(id int64) error
	SearchProducts(filter *request.ProductFilter, pagination common.PaginationRequest) ([]model.Product, int64, error)
	DeleteMultipleProducts(ids []int64) error
	// Example of a transactional operation that involves multiple repository calls
	CreateProductWithDetails(ctx context.Context, product *model.Product, details []model.ProductDetail) error
}
