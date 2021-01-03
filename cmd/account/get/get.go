package get

import (
	"fmt"

	"github.com/sideninja/flow-cli/gateway"
	"github.com/spf13/cobra"
)

// NewCmdGet get account command factory
func NewCmdGet() *cobra.Command {
	var (
		blockHeight int
		filter      string
	)

	getCmd := &cobra.Command{
		Use:   "get <address>",
		Short: "Gets an account by address",
		Long: `
Gets an account by address (address, balance, keys, code)`,
		Args: cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			gw := gateway.CreateGateway("")
			fmt.Println(gw.GetAccount(args[0]))
		},
	}

	getCmd.Flags().IntVarP(&blockHeight, "block-height", "b", -1, "gets an account at the given block height")
	getCmd.Flags().StringVarP(&filter, "filter", "f", "", "filter output to show only provided field (address, balance, code, keys)")

	return getCmd
}

func init() {

}
