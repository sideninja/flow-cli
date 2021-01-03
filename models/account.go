package models

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
	"text/tabwriter"

	"github.com/onflow/flow-go-sdk"
)

// Account is internal structure of the
// account with serialization methods
type Account struct {
	*flow.Account
}

// String serializer for the account
func (a *Account) String(filter string) string {
	switch filter {
	case "":
		return a.stringAll()
	case "balance":
		return fmt.Sprintf("%d", a.Balance)
	case "code":
		return fmt.Sprintf("%s", a.Code)
	case "address":
		return fmt.Sprintf("%s", a.Address)
	case "keys":
		var b bytes.Buffer
		writer := tabwriter.NewWriter(&b, 0, 8, 1, '\t', tabwriter.AlignRight)
		a.stringKeys(writer)
		writer.Flush()
		return b.String()
	default:
		// todo error not valid filter option
		return ""
	}
}

// serialize complete account to string
func (a *Account) stringAll() string {
	// buffer to write to instead of io - eaiser testing and independed from medium
	var b bytes.Buffer
	writer := tabwriter.NewWriter(&b, 0, 8, 1, '\t', tabwriter.AlignRight)

	fmt.Fprintf(writer, "\n")
	fmt.Fprintf(writer, "Address\t %s\n", a.Address)
	fmt.Fprintf(writer, "Balance\t %d\n", a.Balance)

	fmt.Fprintf(writer, "Keys\t %d\n", len(a.Keys))
	a.stringKeys(writer)

	fmt.Fprintf(writer, "\nCode\t\t %s", strings.ReplaceAll(string(a.Code), "\n", "\n\t "))
	fmt.Fprintf(writer, "\n")

	writer.Flush()

	return b.String()
}

func (a *Account) stringKeys(writer *tabwriter.Writer) {
	for i, key := range a.Keys {
		fmt.Fprintf(writer, "\nKey %d\tPublic Key\t %x\n", i, key.PublicKey.Encode())
		fmt.Fprintf(writer, "\tWeight\t %d\n", key.Weight)
		fmt.Fprintf(writer, "\tSignature Algorithm\t %s\n", key.SigAlgo)
		fmt.Fprintf(writer, "\tHash Algorithm\t %s\n", key.HashAlgo)
		fmt.Fprintf(writer, "\n")
	}
}

// JSON serializer for account
func (a *Account) JSON(filter string) string {
	var account []byte
	var err error

	switch filter {
	case "":
		account, err = json.Marshal(a)
	case "balance":
		account, err = json.Marshal(
			struct{ Balance uint64 }{Balance: a.Balance},
		)
	case "code":
		account, err = json.Marshal(
			struct{ Code []byte }{Code: a.Code},
		)
	case "address":
		account, err = json.Marshal(
			struct{ Address flow.Address }{Address: a.Address},
		)
	case "keys":
		account, err = json.Marshal(a.Keys)
	default:
		// todo error not valid filter option
		return ""
	}

	if err != nil {
		fmt.Println("error serializing account")
	}

	return string(account)
}
