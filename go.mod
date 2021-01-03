module github.com/sideninja/flow-cli

go 1.15

replace github.com/golang/protobuf v1.3.5 => google.golang.org/protobuf v1.23.0

require (
	github.com/mitchellh/go-homedir v1.1.0
	github.com/onflow/flow-go-sdk v0.13.1
	github.com/spf13/cobra v1.1.1
	github.com/spf13/viper v1.7.0
	google.golang.org/grpc v1.31.1
)
