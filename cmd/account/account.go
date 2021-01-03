package cmd

import (
	"github.com/sideninja/flow-cli/cmd/account/get"
	"github.com/spf13/cobra"
)

// NewCmdAccount account cmd factory
func NewCmdAccount() *cobra.Command {
	var accountCmd = &cobra.Command{
		Use:     "account",
		Short:   "Get account information, create account, add account key",
		Aliases: []string{"acc"},
		Args:    cobra.ExactArgs(0),
		Long: `
Manage and view flow account.
		`,
	}

	accountCmd.AddCommand(get.NewCmdGet())

	return accountCmd
}

func init() {

}
