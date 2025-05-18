package initialize

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"go-ecommerce-backend-api/global"
	"log"
)

// Configuration
const (
	AWS_REGION = "ap-northeast-1" // Change to your AWS region
)

func InitS3() {
	// Initialize AWS SDK
	cfg, err := config.LoadDefaultConfig(context.TODO(), config.WithRegion(AWS_REGION))
	if err != nil {
		log.Fatalf("unable to load AWS SDK config: %v", err)
	}

	global.Logger.Info("Initializing AWS SDK")
	// Create S3 client
	s3Client := s3.NewFromConfig(cfg)

	global.S3 = s3Client
}
