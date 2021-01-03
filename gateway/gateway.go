package gateway

import (
	"github.com/sideninja/flow-cli/models"
)

// Gateway interface defines all the getter methods
// needed by the cli to be implemented by each gateway
type IGateway interface {
	GetAccount(address string) *models.Account
}
