// shop-api/internal/grpc/auth_client.go
package grpc

import (
	"fmt"
	"sync"

	"store/internal/config"
	"store/internal/grpc/pb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type AuthClient struct {
    client pb.AuthServiceClient
    conn   *grpc.ClientConn
}

var (
    authClientOnce     sync.Once
    authClientInstance *AuthClient
)

func NewAuthClient(addr config.AuthServiceAddress) (*AuthClient, error) {
    var err error
    authClientOnce.Do(func() {
        conn, dialErr := grpc.NewClient(string(addr), 
            grpc.WithTransportCredentials(insecure.NewCredentials()),
        )
        if dialErr != nil {
            err = fmt.Errorf("failed to create auth client: %w", dialErr)
            return
        }
        
        authClientInstance = &AuthClient{
            client: pb.NewAuthServiceClient(conn),
            conn:   conn,
        }
    })
    
    if err != nil {
        return nil, err
    }
    return authClientInstance, nil
}

func (c *AuthClient) Close() error {
    if c.conn != nil {
        return c.conn.Close()
    }
    return nil
}

func (c *AuthClient) GetClient() pb.AuthServiceClient {
    return c.client
}