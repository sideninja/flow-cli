package main

import (
	"fmt"
	"os"

	cmd "github.com/sideninja/flow-cli/cmd"
	"github.com/sideninja/flow-cli/gateway"
)

func main() {
	gateway := gateway.CreateGateway("", "localhost:3569")
	version := "1.0"

	rootCmd := cmd.NewCmdRoot(gateway, version)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
