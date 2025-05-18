package main

import (
	"context"
	"fmt"
	"github.com/gin-contrib/cors"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/swaggo/files"       // swagger embed files
	"github.com/swaggo/gin-swagger" // gin-swagger middleware
	_ "go-ecommerce-backend-api/cmd/swag/docs"
	"go-ecommerce-backend-api/global"
	"go-ecommerce-backend-api/internal/initialize"
)

// Configuration
const (
	PORT       = "8080"
	S3_BUCKET  = "hoapkbucket"
	AWS_REGION = "ap-northeast-1" // Change to your AWS region
	CHUNK_SIZE = 1024 * 1024      // 1MB chunks for streaming
)

// S3Streamer holds the S3 client and bucket name
type S3Streamer struct {
	S3Client   *s3.Client
	BucketName string
}

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
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Content-Type", "Authorization", "X-Requested-With"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	// Initialize AWS SDK
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(AWS_REGION))
	if err != nil {
		log.Fatalf("unable to load AWS SDK config: %v", err)
	}

	global.Logger.Info("Initializing AWS SDK")
	// Create S3 client
	s3Client := s3.NewFromConfig(cfg)

	// Create S3Streamer
	streamer := &S3Streamer{
		S3Client:   s3Client,
		BucketName: S3_BUCKET,
	}

	global.Logger.Info("Initializing AWS SDK Done")

	r.GET("/metrics", gin.WrapH(promhttp.Handler()))
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Add endpoint for streaming video
	r.GET("/stream/:videoKey", streamer.StreamVideo)

	r.Run(":" + strconv.Itoa(global.Config.Server.Port))
}

// StreamVideo handles the video streaming from S3
func (s *S3Streamer) StreamVideo(c *gin.Context) {
	videoKey := c.Param("videoKey")

	// Get the object info from S3
	headResp, err := s.S3Client.HeadObject(c.Request.Context(), &s3.HeadObjectInput{
		Bucket: aws.String(s.BucketName),
		Key:    aws.String(videoKey),
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get object info from S3"})
		return
	}

	fileSize := headResp.ContentLength

	// Parse range header
	rangeHeader := c.GetHeader("Range")
	var start, end int64
	if rangeHeader != "" {
		fmt.Sscanf(rangeHeader, "bytes=%d-%d", &start, &end)
		if end == 0 || end >= *fileSize {
			end = *fileSize - 1
		}
	} else {
		start = 0
		end = *fileSize - 1
	}
	contentLength := end - start + 1

	// Set headers for streaming
	c.Header("Content-Type", *headResp.ContentType)
	c.Header("Accept-Ranges", "bytes")
	c.Header("Content-Length", strconv.FormatInt(contentLength, 10))
	c.Header("Content-Range", fmt.Sprintf("bytes %d-%d/%d", start, end, *fileSize))

	if rangeHeader != "" {
		c.Status(http.StatusPartialContent)
	} else {
		c.Status(http.StatusOK)
	}

	// Get the object from S3 with range
	resp, err := s.S3Client.GetObject(c.Request.Context(), &s3.GetObjectInput{
		Bucket: aws.String(s.BucketName),
		Key:    aws.String(videoKey),
		Range:  aws.String(fmt.Sprintf("bytes=%d-%d", start, end)),
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get object from S3"})
		return
	}
	defer resp.Body.Close()

	// Stream the video
	remaining := contentLength
	buffer := make([]byte, CHUNK_SIZE)
	for remaining > 0 {
		readSize := int64(CHUNK_SIZE)
		if remaining < readSize {
			readSize = remaining
		}

		n, err := io.ReadFull(resp.Body, buffer[:readSize])
		if err != nil && err != io.EOF && err != io.ErrUnexpectedEOF {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error reading from S3"})
			return
		}

		if n > 0 {
			if _, err := c.Writer.Write(buffer[:n]); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Error writing to response"})
				return
			}
			c.Writer.Flush()
			remaining -= int64(n)
		}

		if n == 0 || err == io.EOF {
			break
		}
	}
}
