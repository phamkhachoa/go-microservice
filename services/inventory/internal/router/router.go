package router

//
//import (
//	"fmt"
//	"github.com/gin-gonic/gin"
//	c "go-ecommerce-backend-api/internal/controller"
//	"go-ecommerce-backend-api/internal/middlewares"
//	"net/http"
//)
//
//func AA() gin.HandlerFunc {
//	return func(c *gin.Context) {
//		fmt.Println("Before ->> AA")
//		c.Next()
//		fmt.Println("After ->> AA")
//	}
//}
//
//func BB() gin.HandlerFunc {
//	return func(c *gin.Context) {
//		fmt.Println("Before ->> BB")
//		c.Next()
//		fmt.Println("After ->> BB")
//	}
//}
//
//func CC() gin.HandlerFunc {
//	return func(c *gin.Context) {
//		fmt.Println("Before ->> CC")
//		c.Next()
//		fmt.Println("After ->> CC")
//	}
//}
//
//func NewRouter() *gin.Engine {
//	r := gin.Default()
//	// use the middleware
//	r.Use(middlewares.AuthMiddleware())
//
//	v1 := r.Group("/api/v1")
//	{
//		v1.GET("/ping", Pong)
//		v1.PUT("/ping", Pong)
//		v1.GET("/users/1", c.NewUserController().GetUserById)
//	}
//
//	v2 := r.Group("/api/v2")
//	{
//		v2.GET("/ping", Pong)
//		v2.PUT("/ping", Pong)
//	}
//	return r
//}
//
//func Pong(c *gin.Context) {
//	name := c.DefaultQuery("name", "")
//
//	c.JSONP(http.StatusOK, gin.H{
//		"message": "pong" + name,
//	})
//}
