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
type GRPC struct {
	APIURL string
}

// GetAccount gets account by the address via grpc call
func (g *GRPC) GetAccount(address string) *models.Account {
	// todo move this to singleton for grpc and should be executed by root cmd
	flowClient, err := client.New(g.APIURL, grpc.WithInsecure())

	if err != nil {
		log.Fatal(err)
	}

	// validation of params

	account, err := flowClient.GetAccountAtLatestBlock(
		context.Background(),
		flow.HexToAddress(address),
	)

	if err != nil {
		log.Fatal(err)
	}

	return &models.Account{account}
}
