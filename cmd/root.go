package cmd

import (
	accountCmd "github.com/sideninja/flow-cli/cmd/account"
	"github.com/spf13/cobra"
)

var cfgFile string

// NewCmdRoot root command factory
func NewCmdRoot() *cobra.Command {

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

	rootCmd.AddCommand(accountCmd.NewCmdAccount())

	rootCmd.Flags().BoolP("version", "v", false, "show version information")
	rootCmd.PersistentFlags().BoolVarP(&json, "json", "j", false, "show output format as JSON")

	//rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	return rootCmd
}
