package main

import (
	"fmt"
	"github.com/sideninja/flow-cli/util"

	"github.com/sideninja/flow-cli/cmd"
	"github.com/sideninja/flow-cli/gateway"
	"github.com/sideninja/flow-cli/models"
	"github.com/spf13/viper"
)

func init() {
	util.InitConfig()
}

var version string
var foo string
var boo string

func main() {

	APIURL := viper.GetString("APIURL")

	fmt.Println("HERE", version, foo, boo, "...")

	gateway := gateway.CreateGateway(models.GRPC, APIURL)
	rootCmd := cmd.NewCmdRoot(gateway, version+foo+boo)

	if err := rootCmd.Execute(); err != nil {
		util.HandleError(err)
	}
}
