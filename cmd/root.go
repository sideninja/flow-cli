package cmd

import (
	accountCmd "github.com/sideninja/flow-cli/cmd/account"
	"github.com/sideninja/flow-cli/gateway"
	"github.com/spf13/cobra"
)

var cfgFile string

// NewCmdRoot root command factory
func NewCmdRoot(gateway gateway.IGateway, version string) *cobra.Command {

	var (
		json bool
	)

	var rootCmd = &cobra.Command{
		Use:              "flow",
		Short:            "Flow CLI Tool",
		TraverseChildren: true,
		Long: `
Flow CLI tool to interact with flow emulator.`,
	}

	rootCmd.AddCommand(accountCmd.NewCmdAccount(gateway, version))

	rootCmd.Flags().BoolP("version", "v", false, "show version information")
	rootCmd.PersistentFlags().BoolVarP(&json, "json", "j", false, "show output format as JSON")

	return rootCmd
}
