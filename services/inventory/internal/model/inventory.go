package model

import (
    "time"
)

// Inventory represents the product inventory information
type Inventory struct {
    ID              int64      `json:"id" gorm:"primaryKey;autoIncrement"`
    ProductID       int64      `json:"product_id" gorm:"uniqueIndex;not null"`
    Quantity        int        `json:"quantity" gorm:"not null;default:0"`
    ReservedQuantity int       `json:"reserved_quantity" gorm:"not null;default:0"`
    ReorderPoint    *int       `json:"reorder_point" gorm:"default:null"`
    ReorderQuantity *int       `json:"reorder_quantity" gorm:"default:null"`
    LastRestockDate *time.Time `json:"last_restock_date" gorm:"default:null"`
    UpdatedAt       time.Time  `json:"updated_at" gorm:"autoUpdateTime"`
    CreatedAt       time.Time  `json:"created_at" gorm:"autoCreateTime"`
}

// TableName specifies the table name for the Inventory model
func (Inventory) TableName() string {
    return "product_inventory"
}

// AvailableQuantity returns the quantity available for purchase (total - reserved)
func (i *Inventory) AvailableQuantity() int {
    return i.Quantity - i.ReservedQuantity
}

// IsLowStock checks if the inventory is at or below the reorder point
func (i *Inventory) IsLowStock() bool {
    if i.ReorderPoint == nil {
        return false
    }
    return i.Quantity <= *i.ReorderPoint
}

// Reserve attempts to reserve a specified quantity of the product
// Returns true if successful, false if insufficient quantity
func (i *Inventory) Reserve(quantity int) bool {
    if i.AvailableQuantity() < quantity {
        return false
    }
    
    i.ReservedQuantity += quantity
    return true
}

// Fulfill reduces the inventory by the specified quantity after a successful order
// Returns true if successful, false if insufficient quantity
func (i *Inventory) Fulfill(quantity int) bool {
    if i.ReservedQuantity < quantity {
        return false
    }
    
    i.Quantity -= quantity
    i.ReservedQuantity -= quantity
    return true
}

// Restock increases the inventory quantity and updates the last restock date
func (i *Inventory) Restock(quantity int) {
    i.Quantity += quantity
    now := time.Now()
    i.LastRestockDate = &now
}

// CancelReservation returns reserved items back to available inventory
func (i *Inventory) CancelReservation(quantity int) bool {
    if i.ReservedQuantity < quantity {
        return false
    }
    
    i.ReservedQuantity -= quantity
    return true
}
