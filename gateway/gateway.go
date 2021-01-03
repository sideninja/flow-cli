package gateway

import (
	"github.com/sideninja/flow-cli/models"
)

// IGateway interface defines all the getter methods
// needed by the cli to be implemented by each gateway
type IGateway interface {
	GetAccount(address string) *models.Account
}

// CreateGateway creates a gateway from a factory
func CreateGateway(method string) IGateway {
	// check if testing or rest setting and return test instance
	if method == "rest" {
		return &Rest{}
	}

	return &GRPC{}
}
