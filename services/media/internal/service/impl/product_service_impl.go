package impl

//
//import (
//	"fmt"
//	"github.com/aws/aws-sdk-go-v2/aws"
//	"github.com/aws/aws-sdk-go-v2/service/s3"
//	"github.com/gin-gonic/gin"
//	"go-ecommerce-backend-api/global"
//	"go-ecommerce-backend-api/internal/client"
//	"go-ecommerce-backend-api/internal/model"
//	"go-ecommerce-backend-api/internal/repo"
//	"go-ecommerce-backend-api/internal/service"
//	"io"
//	"net/http"
//	"strconv"
//)
//
//type s3Service struct {
//}
//
//// NewProductService creates a new product service
//func NewS3Service(productRepo repo.ProductRepo, inventoryClient *client.InventoryClient) service.IS3Service {
//	return &s3Service{}
//}
//
//func (s s3Service) StreamFile(bucketName string, fileName string) (*model.Product, error) {
//	//TODO implement me
//	panic("implement me")
//}
//
//func (s *s3Service) StreamFile(c *gin.Context) {
//	videoKey := c.Param("videoKey")
//
//	// Get the object info from S3
//	headResp, err := global.S3.HeadObject(c.Request.Context(), &s3.HeadObjectInput{
//		Bucket: aws.String(s.BucketName),
//		Key:    aws.String(videoKey),
//	})
//	if err != nil {
//		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get object info from S3"})
//		return
//	}
//
//	fileSize := headResp.ContentLength
//
//	// Parse range header
//	rangeHeader := c.GetHeader("Range")
//	var start, end int64
//	if rangeHeader != "" {
//		fmt.Sscanf(rangeHeader, "bytes=%d-%d", &start, &end)
//		if end == 0 || end >= *fileSize {
//			end = *fileSize - 1
//		}
//	} else {
//		start = 0
//		end = *fileSize - 1
//	}
//	contentLength := end - start + 1
//
//	// Set headers for streaming
//	c.Header("Content-Type", *headResp.ContentType)
//	c.Header("Accept-Ranges", "bytes")
//	c.Header("Content-Length", strconv.FormatInt(contentLength, 10))
//	c.Header("Content-Range", fmt.Sprintf("bytes %d-%d/%d", start, end, *fileSize))
//
//	if rangeHeader != "" {
//		c.Status(http.StatusPartialContent)
//	} else {
//		c.Status(http.StatusOK)
//	}
//
//	// Get the object from S3 with range
//	resp, err := s.S3Client.GetObject(c.Request.Context(), &s3.GetObjectInput{
//		Bucket: aws.String(s.BucketName),
//		Key:    aws.String(videoKey),
//		Range:  aws.String(fmt.Sprintf("bytes=%d-%d", start, end)),
//	})
//	if err != nil {
//		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get object from S3"})
//		return
//	}
//	defer resp.Body.Close()
//
//	// Stream the video
//	remaining := contentLength
//	buffer := make([]byte, CHUNK_SIZE)
//	for remaining > 0 {
//		readSize := int64(CHUNK_SIZE)
//		if remaining < readSize {
//			readSize = remaining
//		}
//
//		n, err := io.ReadFull(resp.Body, buffer[:readSize])
//		if err != nil && err != io.EOF && err != io.ErrUnexpectedEOF {
//			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error reading from S3"})
//			return
//		}
//
//		if n > 0 {
//			if _, err := c.Writer.Write(buffer[:n]); err != nil {
//				c.JSON(http.StatusInternalServerError, gin.H{"error": "Error writing to response"})
//				return
//			}
//			c.Writer.Flush()
//			remaining -= int64(n)
//		}
//
//		if n == 0 || err == io.EOF {
//			break
//		}
//	}
//}
