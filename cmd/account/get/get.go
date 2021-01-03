package get

import (
	"fmt"

	"github.com/sideninja/flow-cli/gateway"
	"github.com/spf13/cobra"
)

// NewCmdGet get account command factory
func NewCmdGet(gateway gateway.IGateway, version string) *cobra.Command {
	var (
		blockHeight int
		filter      string
		json        bool
		api         string
	)

	getCmd := &cobra.Command{
		Use:     "get <address>",
		Short:   "Gets an account by address",
		Aliases: []string{"fetch", "g"},
		Long: `
Gets an account by address (address, balance, keys, code)`,
		Args: cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			if api != "" {
				gateway.SetAPIURL(api)
			}

			account := gateway.GetAccount(args[0])

			if !json {
				fmt.Println(account.String(filter))
			} else {
				fmt.Println(account.JSON(filter))
			}
		},
	}

	getCmd.PersistentFlags().BoolVarP(&json, "json", "j", false, "show output format as JSON")
	getCmd.PersistentFlags().StringVarP(&api, "api", "a", "", "set discovery api host on the fly only for this request")
	getCmd.Flags().IntVarP(&blockHeight, "block-height", "b", -1, "gets an account at the given block height")
	getCmd.Flags().StringVarP(&filter, "filter", "f", "", "filter output to show only provided field (address, balance, code, keys)")

	return getCmd
}

func init() {

}
