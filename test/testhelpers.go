package test

import (
	"github.com/onflow/flow-go-sdk"
	"github.com/onflow/flow-go-sdk/crypto"
	"github.com/sideninja/flow-cli/models"
)

// GetMockAccount gets mock account for test
func GetMockAccount() *models.Account {
	sigAlgo := crypto.StringToSignatureAlgorithm("ECDSA_P256")
	pubKey, _ := crypto.DecodePublicKeyHex(sigAlgo, "0a4c9f7bf0c8adf6ebdf4859c11d8e2867e0eaa4af6880e18a0730a0b44a494e976cefa0caf8efb7ec2da469c3f320dab4a2ca72fb340621776f4a1403ae39ed")

	flowAccount := flow.Account{
		Address: flow.HexToAddress("01cf0e2f2f715450"),
		Balance: 10,
		Code:    nil,
		Keys: []*flow.AccountKey{
			&flow.AccountKey{
				Index:          1,
				PublicKey:      pubKey,
				SigAlgo:        sigAlgo,
				HashAlgo:       crypto.HashAlgorithm(1),
				Weight:         1000,
				SequenceNumber: 1,
				Revoked:        false,
			},
		},
	}

	return &models.Account{&flowAccount}
}

// MockGateway for testing
type MockGateway struct{}

// SetAPIURL for testing
func (r *MockGateway) SetAPIURL(url string) {}

// GetAccount for testing
func (r *MockGateway) GetAccount(address string) (*models.Account, error) {
	return GetMockAccount(), nil
}

// GetMockGateway for testing
func GetMockGateway() *MockGateway {
	return &MockGateway{}
}
