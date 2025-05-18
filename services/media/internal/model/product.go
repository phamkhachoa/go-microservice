package model

import (
	"database/sql/driver"
	"encoding/json"
	"errors"
	"time"
)

// JSONMap is a type for handling JSON fields in the database
type JSONMap map[string]interface{}

// Value implements the driver.Valuer interface for JSONMap
func (j JSONMap) Value() (driver.Value, error) {
	if j == nil {
		return nil, nil
	}
	return json.Marshal(j)
}

// Scan implements the sql.Scanner interface for JSONMap
func (j *JSONMap) Scan(value interface{}) error {
	if value == nil {
		*j = nil
		return nil
	}

	bytes, ok := value.([]byte)
	if !ok {
		return errors.New("type assertion to []byte failed")
	}

	return json.Unmarshal(bytes, j)
}

// Product represents the product model
type Product struct {
	ID            int64     `json:"id" gorm:"primaryKey;column:id"`
	Name          string    `json:"name" gorm:"column:name;not null"`
	Description   string    `json:"description" gorm:"column:description"`
	Price         float64   `json:"price" gorm:"column:price;not null"`
	DiscountPrice *float64  `json:"discount_price" gorm:"column:discount_price"`
	Quantity      int       `json:"quantity" gorm:"column:quantity;not null;default:0"`
	CategoryID    *int64    `json:"category_id" gorm:"column:category_id"`
	Thumbnail     string    `json:"thumbnail" gorm:"column:thumbnail"`
	Images        JSONMap   `json:"images" gorm:"column:images;type:json"`
	Attributes    JSONMap   `json:"attributes" gorm:"column:attributes;type:json"`
	Status        int       `json:"status" gorm:"column:status;not null;default:1"`
	UpdatedAt     time.Time `json:"updated_at" gorm:"column:updated_at;autoUpdateTime"`
	CreatedAt     time.Time `json:"created_at" gorm:"column:created_at;autoCreateTime"`
}

// TableName specifies the table name for the Product model
func (Product) TableName() string {
	return "product"
}
