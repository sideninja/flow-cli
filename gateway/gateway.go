package gateway

// Gateway interface defines all the getter methods
// needed by the cli to be implemented by each gateway
type Gateway interface {
	GetAccount() Account
}
