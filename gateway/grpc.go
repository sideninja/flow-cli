package gateway

import (
	"context"
	"log"

	"github.com/sideninja/flow-cli/models"

	"github.com/onflow/flow-go-sdk"
	"github.com/onflow/flow-go-sdk/client"
	"google.golang.org/grpc"
)

// GRPC api implementation
type GRPC struct{}

// GetAccount gets account by the address via grpc call
func (g *GRPC) GetAccount(address string) *models.Account {
	flowClient, err := client.New("localhost:3569", grpc.WithInsecure())

	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()
	addr := flow.HexToAddress(address)
	account, err := flowClient.GetAccountAtLatestBlock(ctx, addr)

	if err != nil {
		log.Fatal(err)
	}

	return &models.Account{account}
}
