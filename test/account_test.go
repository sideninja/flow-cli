package test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var account = GetMockAccount()

func Test_AccountEncapsulation(t *testing.T) {
	assert.Equal(t, account.Address.String(), "01cf0e2f2f715450")
	assert.Equal(t, int(account.Balance), 10)
}

func Test_AccountStringSerialization(t *testing.T) {

	assert.Equal(t, account.String("address"), "01cf0e2f2f715450")
	assert.Equal(t, account.String("balance"), "10")
	assert.Equal(t, account.String("keys"), "\nKey 0\tPublic Key\t\t 0a4c9f7bf0c8adf6ebdf4859c11d8e2867e0eaa4af6880e18a0730a0b44a494e976cefa0caf8efb7ec2da469c3f320dab4a2ca72fb340621776f4a1403ae39ed\n\tWeight\t\t\t 1000\n\tSignature Algorithm\t ECDSA_P256\n\tHash Algorithm\t\t SHA2_256\n\n")

	assert.Equal(t, account.String(""), "\nAddress\t 01cf0e2f2f715450\nBalance\t 10\nKeys\t 1\n\nKey 0\tPublic Key\t\t 0a4c9f7bf0c8adf6ebdf4859c11d8e2867e0eaa4af6880e18a0730a0b44a494e976cefa0caf8efb7ec2da469c3f320dab4a2ca72fb340621776f4a1403ae39ed\n\tWeight\t\t\t 1000\n\tSignature Algorithm\t ECDSA_P256\n\tHash Algorithm\t\t SHA2_256\n\n\nCode\t\t \n")
}

func Test_AccountJSONSerialization(t *testing.T) {
	assert.Equal(t, account.JSON(""), "{\"Address\":\"01cf0e2f2f715450\",\"Balance\":10,\"Code\":null,\"Keys\":[{\"Index\":1,\"PublicKey\":{},\"SigAlgo\":2,\"HashAlgo\":1,\"Weight\":1000,\"SequenceNumber\":1,\"Revoked\":false}],\"Contracts\":null}")
	assert.Equal(t, account.JSON("balance"), "{\"Balance\":10}")
	assert.Equal(t, account.JSON("address"), "{\"Address\":\"01cf0e2f2f715450\"}")
	assert.Equal(t, account.JSON("keys"), "[{\"Index\":1,\"PublicKey\":{},\"SigAlgo\":2,\"HashAlgo\":1,\"Weight\":1000,\"SequenceNumber\":1,\"Revoked\":false}]")
}
