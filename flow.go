package main

import (
	"github.com/sideninja/flow-cli/util"

	"github.com/sideninja/flow-cli/cmd"
	"github.com/sideninja/flow-cli/gateway"
	"github.com/sideninja/flow-cli/models"
	"github.com/spf13/viper"
)

func init() {
	util.InitConfig()
}

func main() {

	APIURL := viper.GetString("APIURL")
	version := "1.0"

	gateway := gateway.CreateGateway(models.GRPC, APIURL)
	rootCmd := cmd.NewCmdRoot(gateway, version)

	if err := rootCmd.Execute(); err != nil {
		util.HandleError(err)
	}
}
