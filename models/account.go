package models

import "github.com/onflow/flow-go-sdk"

// Account is internal structure of the
// account with serialization methods
type Account struct {
	*flow.Account
}

// ToString converts account to string
func (a *Account) ToString() string {
	return a.Address.String()
}
