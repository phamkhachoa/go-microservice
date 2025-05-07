package main

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	_ "go-ecommerce-backend-api/cmd/swag/docs" // which is the generated folder after swag init
	"go-ecommerce-backend-api/global"
	"go-ecommerce-backend-api/internal/initialize"
	"strconv"
)
import "github.com/swaggo/gin-swagger" // gin-swagger middleware
import "github.com/swaggo/files"       // swagger embed files

// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /v1/2024

// @securityDefinitions.basic  BasicAuth

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func main() {
	r := initialize.Run()
	r.GET("/metrics", gin.WrapH(promhttp.Handler()))
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	//r.GET("/swagger/doc.json", func(w http.ResponseWriter, _ *http.Request) {
	//	w.Header().Set("Content-Type", "application/json")
	//	w.WriteHeader(200)
	//	w.Write([]byte(docs.SwaggerInfo.ReadDoc()))
	//})
	r.Run(":" + strconv.Itoa(global.Config.Server.Port))
}
