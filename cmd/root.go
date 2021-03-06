package cmd

import (
	accountCmd "github.com/sideninja/flow-cli/cmd/account"
	configCmd "github.com/sideninja/flow-cli/cmd/config"
	"github.com/sideninja/flow-cli/gateway"
	"github.com/sideninja/flow-cli/util"
	"github.com/spf13/cobra"
)

// NewCmdRoot root command factory
func NewCmdRoot(gateway gateway.IGateway, version string) *cobra.Command {

	var (
		json bool
	)

	var rootCmd = &cobra.Command{
		Use:              "flow",
		Short:            "Flow CLI Tool",
		TraverseChildren: true,
		SilenceErrors:    true,
		SilenceUsage:     true,
		Long: `
Flow CLI tool to interact with flow emulator.`,
	}

	// here we add all commands:
	rootCmd.AddCommand(
		accountCmd.NewCmdAccount(gateway, version),
	)
	rootCmd.AddCommand(
		configCmd.NewCmdConfig(gateway, version),
	)

	// always enabled version and json flags
	rootCmd.Flags().BoolP("version", "v", false, "show version information")
	rootCmd.PersistentFlags().BoolVarP(&json, "json", "j", false, "show output format as JSON")

	// error handling for commands
	rootCmd.SetFlagErrorFunc(func(cmd *cobra.Command, err error) error {
		cmd.Printf("\n\033[31mError: %s\033[0m\n\n", err)
		cmd.Println(cmd.UsageString())
		return util.SilentErr
	})

	return rootCmd
}
