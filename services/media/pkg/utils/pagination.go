package utils

import (
	"fmt"
	"strings"

	"gorm.io/gorm"
)

// ApplyPagination applies pagination to a GORM query
func ApplyPagination(db *gorm.DB, page, limit int64) *gorm.DB {
	offset := (page - 1) * limit
	return db.Offset(int(offset)).Limit(int(limit))
}

// ApplySorting applies sorting to a GORM query
// sortParam format: "field:direction" (e.g., "name:asc", "price:desc")
func ApplySorting(db *gorm.DB, sortParam string) *gorm.DB {
	if sortParam == "" {
		return db
	}

	parts := strings.Split(sortParam, ":")
	if len(parts) != 2 {
		return db
	}

	field := parts[0]
	direction := parts[1]

	// Validate direction
	if direction != "asc" && direction != "desc" {
		direction = "asc"
	}

	// Apply sorting
	orderClause := fmt.Sprintf("%s %s", field, direction)
	return db.Order(orderClause)
}

// GetTotalPages calculates the total number of pages
func GetTotalPages(totalItems, itemsPerPage int64) int64 {
	if itemsPerPage <= 0 {
		return 0
	}

	if totalItems <= 0 {
		return 0
	}

	return (totalItems + itemsPerPage - 1) / itemsPerPage
}
