package gateway

import "github.com/sideninja/flow-cli/models"

// Rest possible implementation
type Rest struct{}

// SetAPIURL sets the url
func (r *Rest) SetAPIURL(url string) {
	// not implemented
}

// GetAccount gets account over rest api
func (r *Rest) GetAccount(address string) *models.Account {
	// not implemented
	return &models.Account{}
}
