package gateway

import (
	"context"

	"github.com/sideninja/flow-cli/models"

	"github.com/onflow/flow-go-sdk"
	"github.com/onflow/flow-go-sdk/client"
	"google.golang.org/grpc"
)

// GRPC api implementation
type GRPC struct {
	APIURL string
}

// SetAPIURL sets the url of the discovery api
func (g *GRPC) SetAPIURL(url string) {
	// todo do url validation

	g.APIURL = url
}

// GetAccount gets account by the address via grpc call
func (g *GRPC) GetAccount(address string) (*models.Account, error) {
	flowClient, err := client.New(g.APIURL, grpc.WithInsecure())

	if err != nil {
		return nil, err
	}

	// validation of params

	account, err := flowClient.GetAccountAtLatestBlock(
		context.Background(),
		flow.HexToAddress(address),
	)

	if err != nil {
		return nil, err
	}

	return &models.Account{account}, nil
}
