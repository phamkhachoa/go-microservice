package initialize

import (
	"github.com/gin-gonic/gin"
	"go-ecommerce-backend-api/global"
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
	//r.Use() // cross
	//r.Use() // limiter
	//productRouter := router.RouterGroupApp.Product

	MainGroup := r.Group("/v1/2024")
	{
		MainGroup.GET("/check-status")
	}
	{
		//productRouter.InitRouter(MainGroup)
	}
	return r
}
