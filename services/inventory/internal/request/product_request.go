package request

// CreateProductRequest represents the request to create a product
type CreateProductRequest struct {
	Name          string                 `json:"name" binding:"required"`
	Description   string                 `json:"description"`
	Price         float64                `json:"price" binding:"required,gt=0"`
	DiscountPrice *float64               `json:"discount_price"`
	Quantity      int                    `json:"quantity" binding:"required,gte=0"`
	CategoryID    *int64                 `json:"category_id"`
	Thumbnail     string                 `json:"thumbnail"`
	Images        map[string]interface{} `json:"images"`
	Attributes    map[string]interface{} `json:"attributes"`
	Status        int                    `json:"status" binding:"required,oneof=1 0 -1"`
}

// UpdateProductRequest represents the request to update a product
type UpdateProductRequest struct {
	Name          string                 `json:"name"`
	Description   string                 `json:"description"`
	Price         *float64               `json:"price" binding:"omitempty,gt=0"`
	DiscountPrice *float64               `json:"discount_price"`
	Quantity      *int                   `json:"quantity" binding:"omitempty,gte=0"`
	CategoryID    *int64                 `json:"category_id"`
	Thumbnail     string                 `json:"thumbnail"`
	Images        map[string]interface{} `json:"images"`
	Attributes    map[string]interface{} `json:"attributes"`
	Status        *int                   `json:"status" binding:"omitempty,oneof=1 0 -1"`
}

// ProductFilter represents the filter for searching products
type ProductFilter struct {
	Name       string   `form:"name"`
	CategoryID *int64   `form:"category_id"`
	MinPrice   *float64 `form:"min_price"`
	MaxPrice   *float64 `form:"max_price"`
	Status     *int     `form:"status"`
}

// DeleteMultipleProductsRequest represents a request to delete multiple products
type DeleteMultipleProductsRequest struct {
	IDs []int64 `json:"ids" binding:"required,min=1"`
}
