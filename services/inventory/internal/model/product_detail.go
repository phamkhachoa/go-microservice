package model

import (
    "time"
)

// ProductDetail represents additional information about a product
type ProductDetail struct {
    ID          int64     `json:"id" gorm:"primaryKey;autoIncrement"`
    ProductID   int64     `json:"product_id" gorm:"index;not null"`
    Type        string    `json:"type" gorm:"size:50;not null"` // e.g., "specification", "feature", "warranty"
    Name        string    `json:"name" gorm:"size:100;not null"`
    Value       string    `json:"value" gorm:"size:255;not null"`
    DisplayOrder int      `json:"display_order" gorm:"default:0"`
    Status      int8      `json:"status" gorm:"default:1"` // 1-active, 0-inactive, -1-deleted
    CreatedAt   time.Time `json:"created_at" gorm:"autoCreateTime"`
    UpdatedAt   time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

// TableName specifies the table name for ProductDetail
func (ProductDetail) TableName() string {
    return "product_detail"
}