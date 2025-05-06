package common

// Paging represents pagination metadata
type Paging struct {
	Page      int `json:"page"`
	PageSize  int `json:"page_size"`
	Total     int `json:"total"`
	TotalPage int `json:"total_page"`
}

// PagedResponse represents a generic paginated API response
type PagedResponse[T any] struct {
	Success bool    `json:"success"`
	Message string  `json:"message,omitempty"`
	Data    []T     `json:"data"`
	Paging  Paging  `json:"paging"`
	Errors  []error `json:"errors,omitempty"`
}

// NewPagedResponse creates a new paginated response
func NewPagedResponse[T any](
	success bool,
	message string,
	data []T,
	page int,
	pageSize int,
	total int,
) PagedResponse[T] {
	return PagedResponse[T]{
		Success: success,
		Message: message,
		Data:    data,
		Paging: Paging{
			Page:      page,
			PageSize:  pageSize,
			Total:     total,
			TotalPage: calculateTotalPages(total, pageSize),
		},
	}
}

func calculateTotalPages(total int, pageSize int) int {
	if pageSize <= 0 {
		return 0
	}
	return (total + pageSize - 1) / pageSize
}
