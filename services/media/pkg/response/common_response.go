package response

import (
	"go-ecommerce-backend-api/pkg/i18n"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// Response codes
const (
	CodeSuccess            = 200
	CodeBadRequest         = 400
	CodeUnauthorized       = 401
	CodeForbidden          = 403
	CodeNotFound           = 404
	CodeInternalError      = 500
	CodeServiceUnavailable = 503
)

// Standard messages for response codes
var standardMessages = map[int]string{
	CodeSuccess:            "Success",
	CodeBadRequest:         "Bad request",
	CodeUnauthorized:       "Unauthorized",
	CodeForbidden:          "Forbidden",
	CodeNotFound:           "Not found",
	CodeInternalError:      "Internal server error",
	CodeServiceUnavailable: "Service unavailable",
}

func getTranslatedMessage(c *gin.Context, messageCode string) string {
	lang := c.GetHeader("Accept-Language")
	if lang == "" {
		lang = "vi" // Default to Vietnamese
	}
	return i18n.Translate(lang, messageCode, nil)
}

// Response represents a standard API response structure
type Response struct {
	Code      int         `json:"code"`
	Message   string      `json:"message"`
	Data      interface{} `json:"data,omitempty"`
	Timestamp time.Time   `json:"timestamp"`
}

// PaginationMeta contains metadata for paginated responses
type PaginationMeta struct {
	Page       int64 `json:"page"`
	Limit      int64 `json:"limit"`
	TotalItems int64 `json:"total_items"`
	TotalPages int64 `json:"total_pages"`
}

// PaginatedResponse represents a paginated API response
type PaginatedResponse struct {
	Code      int            `json:"code"`
	Message   string         `json:"message"`
	Data      interface{}    `json:"data,omitempty"`
	Meta      PaginationMeta `json:"meta"`
	Timestamp time.Time      `json:"timestamp"`
}

// NewResponse creates a new standard response
func NewResponse(code int, message string, data interface{}) Response {
	if message == "" {
		if msg, exists := standardMessages[code]; exists {
			message = msg
		}
	}

	return Response{
		Code:      code,
		Message:   message,
		Data:      data,
		Timestamp: time.Now(),
	}
}

// NewPaginatedResponse creates a new paginated response
func NewPaginatedResponse(code int, message string, data interface{}, page, limit, totalItems int64) PaginatedResponse {
	if message == "" {
		if msg, exists := standardMessages[code]; exists {
			message = msg
		}
	}

	var totalPages int64 = 0
	if limit > 0 {
		totalPages = (totalItems + limit - 1) / limit
	}

	return PaginatedResponse{
		Code:    code,
		Message: message,
		Data:    data,
		Meta: PaginationMeta{
			Page:       page,
			Limit:      limit,
			TotalItems: totalItems,
			TotalPages: totalPages,
		},
		Timestamp: time.Now(),
	}
}

// Success sends a success response
func Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, NewResponse(CodeSuccess, "", data))
}

// SuccessWithMessage sends a success response with a custom message
func SuccessWithMessage(c *gin.Context, message string, data interface{}) {
	c.JSON(http.StatusOK, NewResponse(CodeSuccess, message, data))
}

// Created sends a created response
func Created(c *gin.Context, data interface{}) {
	c.JSON(http.StatusCreated, NewResponse(CodeSuccess, "Resource created successfully", data))
}

// Paginated sends a paginated response
func Paginated(c *gin.Context, data interface{}, page, limit, totalItems int64) {
	c.JSON(http.StatusOK, NewPaginatedResponse(CodeSuccess, "", data, page, limit, totalItems))
}

// PaginatedWithMessage sends a paginated response with a custom message
func PaginatedWithMessage(c *gin.Context, message string, data interface{}, page, limit, totalItems int64) {
	c.JSON(http.StatusOK, NewPaginatedResponse(CodeSuccess, message, data, page, limit, totalItems))
}

// BadRequest sends a bad request error response
func BadRequest(c *gin.Context, messageCode string) {
	response := NewResponse(CodeBadRequest, messageCode, nil)
	response.Message = getTranslatedMessage(c, messageCode)
	c.JSON(http.StatusBadRequest, response)
}

// Unauthorized sends an unauthorized error response
// Unauthorized sends an unauthorized error response
func Unauthorized(c *gin.Context, messageCode string) {
	if messageCode == "" {
		messageCode = "unauthorized"
	}
	response := NewResponse(CodeUnauthorized, messageCode, nil)
	response.Message = getTranslatedMessage(c, messageCode)
	c.JSON(http.StatusUnauthorized, response)
}

// Forbidden sends a forbidden error response
func Forbidden(c *gin.Context, message string) {
	if message == "" {
		message = standardMessages[CodeForbidden]
	}
	c.JSON(http.StatusForbidden, NewResponse(CodeForbidden, message, nil))
}

// NotFound sends a not found error response
func NotFound(c *gin.Context, message string) {
	if message == "" {
		message = standardMessages[CodeNotFound]
	}
	c.JSON(http.StatusNotFound, NewResponse(CodeNotFound, message, nil))
}

// InternalError sends an internal server error response
func InternalError(c *gin.Context, message string) {
	if message == "" {
		message = standardMessages[CodeInternalError]
	}
	c.JSON(http.StatusInternalServerError, NewResponse(CodeInternalError, message, nil))
}

// ServiceUnavailable sends a service unavailable error response
func ServiceUnavailable(c *gin.Context, message string) {
	if message == "" {
		message = standardMessages[CodeServiceUnavailable]
	}
	c.JSON(http.StatusServiceUnavailable, NewResponse(CodeServiceUnavailable, message, nil))
}

// Error sends an appropriate error response based on the HTTP status code
func Error(c *gin.Context, statusCode int, message string) {
	switch statusCode {
	case http.StatusBadRequest:
		BadRequest(c, message)
	case http.StatusUnauthorized:
		Unauthorized(c, message)
	case http.StatusForbidden:
		Forbidden(c, message)
	case http.StatusNotFound:
		NotFound(c, message)
	case http.StatusServiceUnavailable:
		ServiceUnavailable(c, message)
	default:
		InternalError(c, message)
	}
}
