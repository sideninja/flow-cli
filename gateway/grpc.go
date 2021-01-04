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
func (g *GRPC) GetAccount(addr string) (*models.Account, error) {
	flowClient, err := client.New(g.APIURL, grpc.WithInsecure())

	if err != nil {
		return nil, err
	}

	chainID := flow.ChainID("flow-emulator") // todo refactor this to config if used on other networks
	address := flow.HexToAddress(addr)

	if !address.IsValid(chainID) {
		return nil, &InvalidAddress{}
	}

	account, err := flowClient.GetAccountAtLatestBlock(
		context.Background(),
		address,
	)

	if err != nil {
		return nil, err
	}

	return &models.Account{account}, nil
}

// InvalidAddress error for wrong account address
type InvalidAddress struct{}

func (m *InvalidAddress) Error() string {
	return "Invalid Address"
}
