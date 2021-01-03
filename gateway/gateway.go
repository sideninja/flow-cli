package gateway

import (
	"github.com/sideninja/flow-cli/models"
)

// IGateway interface defines all the getter methods
// needed by the cli to be implemented by each gateway
type IGateway interface {
	GetAccount(address string) (*models.Account, error)
	SetAPIURL(url string)
}

// CreateGateway creates a gateway from a factory
func CreateGateway(method string, url string) IGateway {
	// check if testing or rest setting and return test instance
	if method == models.REST {
		return &Rest{}
	}

	return &GRPC{
		APIURL: url,
	}
}
