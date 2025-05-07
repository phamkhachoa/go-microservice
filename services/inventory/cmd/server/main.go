package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
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
	// Initialize Gin router
	r := initialize.Run()

	// Configure Gin routes
	r.GET("/metrics", gin.WrapH(promhttp.Handler()))
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Start HTTP server in a goroutine
	go func() {
		httpPort := strconv.Itoa(global.Config.Server.Port)
		global.Logger.Info("Starting HTTP server", zap.String("port", httpPort))
		if err := r.Run(":" + httpPort); err != nil {
			global.Logger.Error("Failed to start HTTP server", zap.Error(err))
		}
	}()

	// Get gRPC port from environment or use default
	grpcPort := os.Getenv("GRPC_PORT")
	if grpcPort == "" {
		grpcPort = "50051"
	}

	// Create gRPC listener
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", grpcPort))
	if err != nil {
		global.Logger.Error("Failed to listen", zap.Error(err), zap.String("port", grpcPort))
		return
	}

	// Create gRPC server
	server := grpcServer.NewServer()

	// Register our gRPC services
	inventoryServer, _ := wire.InitInventoryGrpcServer()
	inventoryPb.RegisterInventoryServiceServer(server, inventoryServer)

	// Start gRPC server in a goroutine
	go func() {
		global.Logger.Info("Starting gRPC server", zap.String("port", grpcPort))
		if err := server.Serve(lis); err != nil {
			global.Logger.Error("Failed to serve gRPC", zap.Error(err))
		}
	}()

	// Wait for interrupt signal to gracefully shut down both servers
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	global.Logger.Info("Shutting down servers...")
	server.GracefulStop()
	global.Logger.Info("gRPC server stopped")
	// Note: Gin doesn't have a built-in graceful shutdown, but you could implement one if needed
	global.Logger.Info("HTTP server stopped")
}
