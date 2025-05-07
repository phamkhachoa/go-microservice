package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go-ecommerce-backend-api/global"
	"go-ecommerce-backend-api/internal/initialize"
	"go-ecommerce-backend-api/internal/wire"
	inventoryPb "go-ecommerce-backend-api/proto"
	"go.uber.org/zap"
	grpcServer "google.golang.org/grpc"
	"net"
	"os"
	"os/signal"
	"strconv"
	"syscall"
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

// @host      localhost:8002
// @BasePath  /v1/2024

// @securityDefinitions.basic  BasicAuth

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func main() {
	r := initialize.Run()

	// grpc
	// Get port from environment or use default
	port := os.Getenv("GRPC_PORT")
	if port == "" {
		port = "50051"
	}

	// Create listener
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		global.Logger.Error("Failed to listen", zap.Error(err), zap.String("port", port))
	}

	// Create gRPC server
	server := grpcServer.NewServer()

	// register our grpc services
	inventoryServer, _ := wire.InitInventoryGrpcServer()
	// register the OrderServiceServer
	inventoryPb.RegisterInventoryServiceServer(server, inventoryServer)

	// Start server in a goroutine
	go func() {
		// Start server in a goroutine
		global.Logger.Info("Starting gRPC server", zap.String("port", port))
		if err := server.Serve(lis); err != nil {
			global.Logger.Error("Failed to serve", zap.Error(err))
		}
	}()

	// Wait for interrupt signal to gracefully shut down the server
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	global.Logger.Info("Shutting down gRPC server...")
	server.GracefulStop()
	global.Logger.Info("Server stopped")

	r.GET("/metrics", gin.WrapH(promhttp.Handler()))
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run(":" + strconv.Itoa(global.Config.Server.Port))
}
