package util

import (
	"errors"
	"fmt"
	"os"

	"github.com/manifoldco/promptui"
	"github.com/onflow/flow-go-sdk/client"
	"github.com/spf13/viper"
	"google.golang.org/grpc/codes"
)

// SilentErr is silent error passed through
var SilentErr = errors.New("SilentErr")

// PromptURL enters discovery api url if missing
func PromptURL() {

	validate := func(input string) error {
		/*_, err := strconv.ParseFloat(input, 64)
		if err != nil {
			return errors.New("Invalid url")
		}*/
		return nil
	}

	prompt := promptui.Prompt{
		Label:    "Enter Flow emulator discovery API URL",
		Validate: validate,
	}

	result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}

	viper.Set("APIURL", result)
	viper.WriteConfigAs("./config.yaml")
}

// HandleError main error handler for commands
func HandleError(err error) {
	if err != SilentErr {
		// check if error is grpc error
		rpcError := client.RPCError{}
		if errors.As(err, &rpcError) {
			handleGrpcError(rpcError)
		} else { // unhandeled error - can be improved
			fmt.Printf("\n\033[31mUnhandeled Error: %s\033[0m\n\n", err)
		}
	}

	os.Exit(1)
}

func handleGrpcError(rpcError client.RPCError) {
	switch rpcError.GRPCStatus().Code() {
	case codes.InvalidArgument:
		fmt.Printf("\n\033[31mInvalid Arguments: %s\033[0m\n\n", rpcError.GRPCStatus().Message())
	case codes.Unavailable:
		fmt.Printf("\nConnection to Flow Emulator was not successful, please make sure your api url is correct.\n\nSet it by using: flow config \n")
		fmt.Printf("\n\033[31mConnection Error: %s\033[0m\n\n", rpcError.GRPCStatus().Message())
		PromptURL()
	default:
		fmt.Printf("\n\033[31mUnhandeled Error: %s\033[0m\n\n", rpcError.GRPCStatus().Err())
	}
}

// InitConfig initialize config for viper
func InitConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.SetDefault("APIURL", "localhost:3569")

	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Println("ERROR not found")
		} else {
			// Config file was found but another error was produced
		}
	}
}
