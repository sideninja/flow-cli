package gateway

import "github.com/sideninja/flow-cli/models"

// Rest possible implementation
type Rest struct{}

// GetAccount gets account over rest api
func (r *Rest) GetAccount(address string) *models.Account {
	// not implemented
	return &models.Account{}
}
