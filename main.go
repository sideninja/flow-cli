package main

import (
	"fmt"
	"os"

	cmd "github.com/sideninja/flow-cli/cmd"
	"github.com/sideninja/flow-cli/gateway"
	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.SetDefault("APIURL", "localhost:3569")

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {

		} else {
			// Config file was found but another error was produced
		}
	}
}

func main() {

	APIURL := viper.GetString("APIURL")

	gateway := gateway.CreateGateway("", APIURL)
	version := "1.0"

	rootCmd := cmd.NewCmdRoot(gateway, version)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
