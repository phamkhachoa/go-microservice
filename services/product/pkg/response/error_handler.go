package response

import (
	"errors"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

// ErrorHandler handles common errors and returns appropriate responses
func ErrorHandler(c *gin.Context, err error) {
	// Handle validation errors
	var validationErrors validator.ValidationErrors
	if errors.As(err, &validationErrors) {
		messages := make([]string, 0, len(validationErrors))
		for _, e := range validationErrors {
			messages = append(messages, formatValidationError(e))
		}
		BadRequest(c, strings.Join(messages, "; "))
		return
	}

	// Handle GORM errors
	if errors.Is(err, gorm.ErrRecordNotFound) {
		NotFound(c, "Resource not found")
		return
	}

	// Handle other specific errors
	switch {
	case strings.Contains(err.Error(), "duplicate"):
		BadRequest(c, "Duplicate entry detected")
	case strings.Contains(err.Error(), "unauthorized"):
		Unauthorized(c, err.Error())
	case strings.Contains(err.Error(), "forbidden"):
		Forbidden(c, err.Error())
	case strings.Contains(err.Error(), "not found"):
		NotFound(c, err.Error())
	default:
		InternalError(c, err.Error())
	}
}

// formatValidationError formats a validation error into a readable message
func formatValidationError(e validator.FieldError) string {
	field := e.Field()

	switch e.Tag() {
	case "required":
		return field + " is required"
	case "email":
		return field + " must be a valid email address"
	case "min":
		return field + " must be at least " + e.Param() + " characters long"
	case "max":
		return field + " must be at most " + e.Param() + " characters long"
	case "gt":
		return field + " must be greater than " + e.Param()
	case "gte":
		return field + " must be greater than or equal to " + e.Param()
	case "lt":
		return field + " must be less than " + e.Param()
	case "lte":
		return field + " must be less than or equal to " + e.Param()
	case "oneof":
		return field + " must be one of [" + e.Param() + "]"
	default:
		return field + " failed validation: " + e.Tag()
	}
}

// HandleBindingError handles errors from binding requests
func HandleBindingError(c *gin.Context, err error) {
	ErrorHandler(c, err)
}
