package controller

import (
	"fmt"
	"go-ecommerce-backend-api/internal/common"
	"go-ecommerce-backend-api/internal/request"
	"go-ecommerce-backend-api/internal/service"
	"go-ecommerce-backend-api/pkg/response"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ProductController struct {
	productService service.IProductService
}

// NewProductController creates a new product controller
func NewProductController(productService service.IProductService) *ProductController {
	return &ProductController{
		productService: productService,
	}
}

// CreateProduct godoc
// @Summary Create a new product
// @Description Create a new product with the provided details
// @Tags products
// @Accept json
// @Produce json
// @Param product body request.CreateProductRequest true "Product details"
// @Success 201 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /api/v1/products [post]
func (c *ProductController) CreateProduct(ctx *gin.Context) {
	var req request.CreateProductRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.HandleBindingError(ctx, err)
		return
	}

	product, err := c.productService.CreateProduct(&req)
	if err != nil {
		response.ErrorHandler(ctx, err)
		return
	}

	response.Created(ctx, product)
}

// GetProduct godoc
// @Summary Get a product by ID
// @Description Get a product by its ID
// @Tags products
// @Accept json
// @Produce json
// @Param id path int true "Product ID"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 404 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /api/v1/products/{id} [get]
func (c *ProductController) GetProduct(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		response.BadRequest(ctx, "Invalid product ID")
		return
	}

	product, err := c.productService.GetProductByID(id)
	if err != nil {
		response.ErrorHandler(ctx, err)
		return
	}

	response.SuccessWithMessage(ctx, "Product retrieved successfully", product)
}

// UpdateProduct godoc
// @Summary Update a product
// @Description Update a product with the provided details
// @Tags products
// @Accept json
// @Produce json
// @Param id path int true "Product ID"
// @Param product body request.UpdateProductRequest true "Product details to update"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 404 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /api/v1/products/{id} [put]
func (c *ProductController) UpdateProduct(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		response.BadRequest(ctx, "Invalid product ID")
		return
	}

	var req request.UpdateProductRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.HandleBindingError(ctx, err)
		return
	}

	product, err := c.productService.UpdateProduct(id, &req)
	if err != nil {
		response.ErrorHandler(ctx, err)
		return
	}

	response.SuccessWithMessage(ctx, "Product updated successfully", product)
}

// DeleteProduct godoc
// @Summary Delete a product
// @Description Delete a product by its ID
// @Tags products
// @Accept json
// @Produce json
// @Param id path int true "Product ID"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 404 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /api/v1/products/{id} [delete]
func (c *ProductController) DeleteProduct(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		response.BadRequest(ctx, "Invalid product ID")
		return
	}

	err = c.productService.DeleteProduct(id)
	if err != nil {
		response.ErrorHandler(ctx, err)
		return
	}

	response.SuccessWithMessage(ctx, "Product deleted successfully", nil)
}

// SearchProducts godoc
// @Summary Search for products
// @Description Search for products based on various filters
// @Tags products
// @Accept json
// @Produce json
// @Param name query string false "Product name"
// @Param category_id query int false "Category ID"
// @Param min_price query number false "Minimum price"
// @Param max_price query number false "Maximum price"
// @Param status query int false "Product status"
// @Param page query int false "Page number"
// @Param limit query int false "Items per page"
// @Param sort query string false "Sort field and direction (e.g., name:asc, price:desc)"
// @Success 200 {object} response.PaginatedResponse
// @Failure 400 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /api/v1/products/search [get]
func (c *ProductController) SearchProducts(ctx *gin.Context) {
	var filter request.ProductFilter
	if err := ctx.ShouldBindQuery(&filter); err != nil {
		response.HandleBindingError(ctx, err)
		return
	}

	// Get pagination parameters from query
	pagination := common.GetPaginationFromQuery(
		ctx.DefaultQuery("page", "1"),
		ctx.DefaultQuery("limit", "10"),
		ctx.DefaultQuery("sort", "id:desc"),
	)

	// Call service with filter and pagination
	result, totalItems, err := c.productService.SearchProducts(&filter, pagination)
	if err != nil {
		response.ErrorHandler(ctx, err)
		return
	}

	response.PaginatedWithMessage(
		ctx,
		"Products retrieved successfully",
		result,
		pagination.Page,
		pagination.Limit,
		totalItems,
	)
}

// DeleteMultipleProducts godoc
// @Summary Delete multiple products
// @Description Delete multiple products by their IDs
// @Tags products
// @Accept json
// @Produce json
// @Param request body request.DeleteMultipleProductsRequest true "Product IDs to delete"
// @Success 200 {object} response.Response
// @Failure 400 {object} response.Response
// @Failure 404 {object} response.Response
// @Failure 500 {object} response.Response
// @Router /api/v1/products/batch-delete [post]
func (c *ProductController) DeleteMultipleProducts(ctx *gin.Context) {
	var req request.DeleteMultipleProductsRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		response.HandleBindingError(ctx, err)
		return
	}

	// Validate request
	if len(req.IDs) == 0 {
		response.BadRequest(ctx, "No product IDs provided")
		return
	}

	// Call service to delete products
	err := c.productService.DeleteMultipleProducts(req.IDs)
	if err != nil {
		response.ErrorHandler(ctx, err)
		return
	}

	response.SuccessWithMessage(ctx, fmt.Sprintf("Successfully deleted %d products", len(req.IDs)), nil)
}
