package config

import (
	"github.com/sideninja/flow-cli/gateway"
	"github.com/sideninja/flow-cli/util"
	"github.com/spf13/cobra"
)

// NewCmdConfig configuration command
func NewCmdConfig(gateway gateway.IGateway, version string) *cobra.Command {

	var configCmd = &cobra.Command{
		Use:   "config <param>",
		Short: "Save persistent configuration",
		Long: `
Save persistent configuration variables.
		`,
		Args: cobra.ExactArgs(0),
		RunE: func(c *cobra.Command, args []string) error {
			util.PromptURL()
			return nil
		},
	}

	return configCmd
}
