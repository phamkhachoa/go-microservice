package common

import "strconv"

// PaginationRequest represents common pagination parameters
type PaginationRequest struct {
	Page  int64  `form:"page" json:"page"`
	Limit int64  `form:"limit" json:"limit"`
	Sort  string `form:"sort" json:"sort"`
}

// GetDefaultPagination returns a pagination request with default values
func GetDefaultPagination() PaginationRequest {
	return PaginationRequest{
		Page:  1,
		Limit: 10,
		Sort:  "id:desc",
	}
}

// GetPaginationFromQuery extracts pagination parameters from query parameters
func GetPaginationFromQuery(page, limit, sort string) PaginationRequest {
	pagination := GetDefaultPagination()

	if pageVal, err := strconv.ParseInt(page, 10, 64); err == nil && pageVal > 0 {
		pagination.Page = pageVal
	}

	if limitVal, err := strconv.ParseInt(limit, 10, 64); err == nil && limitVal > 0 {
		pagination.Limit = limitVal
	}

	if sort != "" {
		pagination.Sort = sort
	}

	return pagination
}
