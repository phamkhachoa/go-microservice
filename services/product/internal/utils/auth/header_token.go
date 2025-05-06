package auth

import (
	"github.com/gin-gonic/gin"
	"strings"
)

func ExtractBearerToken(c *gin.Context) (string, bool) {
	authHeader := c.Request.Header.Get("Authorization")

	if strings.HasPrefix(authHeader, "Bearer ") || strings.HasPrefix(authHeader, "bearer ") {
		return strings.TrimPrefix(authHeader, "Bearer "), true
	}

	return "", false
}
