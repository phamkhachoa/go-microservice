package client

import (
	"context"
	"fmt"
	"go-ecommerce-backend-api/global"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	inventorypb "go-ecommerce-backend-api/proto"
)

// InventoryClient is a client for the inventory service
type InventoryClient struct {
	client inventorypb.InventoryServiceClient
	conn   *grpc.ClientConn
}

// NewInventoryClient creates a new inventory client
func NewInventoryClient() (*InventoryClient, error) {
	inventoryServiceAddr := global.Config.GrpcServer.InventoryGrpcServer
	// Replace with the actual address of your inventory service
	conn, err := grpc.Dial(
		inventoryServiceAddr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithBlock(),
	)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to inventory service: %w", err)
	}

	client := inventorypb.NewInventoryServiceClient(conn)
	return &InventoryClient{
		client: client,
		conn:   conn,
	}, nil
}

// Close closes the connection to the inventory service
func (c *InventoryClient) Close() error {
	if c.conn != nil {
		return c.conn.Close()
	}
	return nil
}

// GetProductInventory gets inventory information for a product
func (c *InventoryClient) GetProductInventory(ctx context.Context, productID int64) (*inventorypb.InventoryResponse, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	resp, err := c.client.GetInventoryByProductID(ctx, &inventorypb.GetInventoryRequest{
		ProductId: productID,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to get product inventory: %w", err)
	}

	return resp, nil
}
