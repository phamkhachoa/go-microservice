package middlewares

import (
	"context"
	"github.com/gin-gonic/gin"
	"go-ecommerce-backend-api/global"
	"go-ecommerce-backend-api/internal/utils/auth"
	"go.uber.org/zap"
	"log"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// get the request url path
		uri := c.Request.URL.Path
		global.Logger.Info("uri", zap.String("uri", uri))
		jwtToken, err := auth.ExtractBearerToken(c)

		if !err {
			c.AbortWithStatusJSON(401, gin.H{"code": 40001, "err": "Unauthorized"})
			return
		}

		// validate jwt
		claims, error := auth.VerifyTokenSubject(jwtToken)
		if error != nil {
			c.AbortWithStatusJSON(401, gin.H{"code": 40001, "err": "Unauthorized"})
			return
		}

		// update claims to context
		log.Println("claims::: UUID::", claims.Subject) // 11clitoken....
		ctx := context.WithValue(c.Request.Context(), "subjectUUID", claims.Subject)
		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}
