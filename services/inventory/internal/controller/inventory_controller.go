package controller

import (
	"go-ecommerce-backend-api/internal/service"
	"go-ecommerce-backend-api/pkg/response"
	"strconv"

	"github.com/gin-gonic/gin"
)

type InventoryController struct {
	inventoryService service.IInventoryService
}

// NewProductController creates a new product controller
func NewInventoryController(inventoryService service.IInventoryService) *InventoryController {
	return &InventoryController{
		inventoryService: inventoryService,
	}
}

// GetByProductId godoc
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
func (c *InventoryController) GetByProductId(ctx *gin.Context) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		response.BadRequest(ctx, "Invalid product ID")
		return
	}

	product, err := c.inventoryService.GetByProductId(id)
	if err != nil {
		response.ErrorHandler(ctx, err)
		return
	}

	response.SuccessWithMessage(ctx, "Product retrieved successfully", product)
}
