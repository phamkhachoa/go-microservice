package initialize

import (
	"github.com/gin-contrib/cors"
	"go-ecommerce-backend-api/global"
	"go-ecommerce-backend-api/internal/router"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	//r := gin.Default()

	var r *gin.Engine
	if global.Config.Server.Mode == "dev" {
		gin.SetMode(gin.DebugMode)
		gin.ForceConsoleColor()
		r = gin.Default()
	} else {
		gin.SetMode(gin.ReleaseMode)
		r = gin.New()
	}

	// middlewares
	//r.Use() // logging
	// Use the CORS middleware
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type", "Authorization", "X-Requested-With"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	//r.Use() // limiter
	inventoryRouter := router.RouterGroupApp.Inventory

	MainGroup := r.Group("/api/inventory/v0")
	{
		MainGroup.GET("/check-status")
		MainGroup.GET("/ping", Pong)
	}
	{
		inventoryRouter.InitRouter(MainGroup)
	}
	return r
}

func Pong(c *gin.Context) {
	c.JSONP(http.StatusOK, gin.H{
		"message": "pong",
	})
}

// CORS middleware to add headers to every response
func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Set CORS headers
		w.Header().Set("Access-Control-Allow-Origin", "*") // Allow any origin
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization, X-Requested-With")

		// Handle preflight OPTIONS requests
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}

		// Process the next handler
		next.ServeHTTP(w, r)
	})
}
