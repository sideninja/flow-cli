package models

import (
	"fmt"

	"github.com/onflow/flow-go-sdk"
)

// Account is internal structure of the
// account with serialization methods
type Account struct {
	*flow.Account
}

// String serializer for the account
func (a *Account) String() string {
	return fmt.Sprintf("address: %s, balance: %d", a.Address.String(), a.Balance)
}
