package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/onflow/flow-go-sdk/client"
	"github.com/sideninja/flow-cli/cmd"
	cliCmd "github.com/sideninja/flow-cli/cmd"
	"github.com/sideninja/flow-cli/gateway"
	"github.com/sideninja/flow-cli/models"
	"github.com/spf13/viper"
	"google.golang.org/grpc/codes"
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
	version := "1.0"

	gateway := gateway.CreateGateway(models.GRPC, APIURL)

	rootCmd := cliCmd.NewCmdRoot(gateway, version)

	if err := rootCmd.Execute(); err != nil {
		handleError(err)
		os.Exit(1)
	}
}

func handleError(err error) {
	if err != cmd.SilentErr {
		// check if error is grpc error
		rpcError := client.RPCError{}
		if errors.As(err, &rpcError) {
			handleGrpcError(rpcError)
		} else { // unhandeled error - can be improved
			fmt.Printf("\n\033[31mUnhandeled Error: %s\033[0m\n\n", err)
		}
	}
}

func handleGrpcError(rpcError client.RPCError) {
	switch rpcError.GRPCStatus().Code() {
	case codes.InvalidArgument:
		fmt.Printf("\n\033[31mInvalid Arguments: %s\033[0m\n\n", rpcError.GRPCStatus().Message())
	case codes.Unavailable:
		fmt.Printf("\nConnection to Flow Emulator was not successful, please make sure your api url is correct. You can set it by: flow [COMMAND] --api URL \n")
		fmt.Printf("\n\033[31mConnection Error: %s\033[0m\n\n", rpcError.GRPCStatus().Message())
	default:
		fmt.Printf("\n\033[31mUnhandeled Error: %s\033[0m\n\n", rpcError.GRPCStatus().Err())
	}
}
