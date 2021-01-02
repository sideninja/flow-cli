module gregor.net/cli

go 1.15

replace github.com/golang/protobuf v1.3.5 => google.golang.org/protobuf v1.23.0

require (
	github.com/onflow/flow-go-sdk v0.13.1
	github.com/spf13/cobra v1.1.1 // indirect
	google.golang.org/grpc v1.31.1
)
