package test

import (
	"bytes"
	"io/ioutil"
	"testing"

	"github.com/sideninja/flow-cli/cmd/account/get"

	"github.com/sideninja/flow-cli/cmd"
	"github.com/sideninja/flow-cli/gateway"
	"github.com/stretchr/testify/assert"
)

func Test_RootCommand(t *testing.T) {
	var gateway gateway.IGateway = GetMockGateway()
	b := bytes.NewBufferString("")

	command := cmd.NewCmdRoot(gateway, "1.0")
	command.SetOut(b)
	command.Execute()

	out, err := ioutil.ReadAll(b)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, string(out), "\nFlow CLI tool to interact with flow emulator.\n\nUsage:\n  flow [command]\n\nAvailable Commands:\n  account     Get account information, create account, add account key\n  config      Save persistent configuration\n  help        Help about any command\n\nFlags:\n  -h, --help      help for flow\n  -j, --json      show output format as JSON\n  -v, --version   show version information\n\nUse \"flow [command] --help\" for more information about a command.\n")
}

func Test_AccountCommandGeneral(t *testing.T) {
	var gateway gateway.IGateway = GetMockGateway()
	b := bytes.NewBufferString("")

	command := get.NewCmdGet(gateway, "1.0")
	command.SetOut(b)
	command.Execute()

	out, err := ioutil.ReadAll(b)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, string(out), "Usage:\n  get <address> [flags]\n\nAliases:\n  get, fetch, g\n\nFlags:\n  -a, --api string         set discovery api host on the fly only for this request\n  -b, --block-height int   gets an account at the given block height (default -1)\n  -f, --filter string      filter output to show only provided field (address, balance, code, keys)\n  -h, --help               help for get\n  -j, --json               show output format as JSON\n\n")
}

func Test_AccountTooManyArgs(t *testing.T) {
	var gateway gateway.IGateway = GetMockGateway()
	b := bytes.NewBufferString("")

	command := get.NewCmdGet(gateway, "1.0")

	command.SetArgs([]string{"123", "123"})
	command.SetOut(b)
	command.SetErr(b)
	command.Execute()

	out, err := ioutil.ReadAll(b)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, string(out), "Error: accepts 1 arg(s), received 2\nUsage:\n  get <address> [flags]\n\nAliases:\n  get, fetch, g\n\nFlags:\n  -a, --api string         set discovery api host on the fly only for this request\n  -b, --block-height int   gets an account at the given block height (default -1)\n  -f, --filter string      filter output to show only provided field (address, balance, code, keys)\n  -h, --help               help for get\n  -j, --json               show output format as JSON\n\n")
}

func Test_AccountWrongArgs(t *testing.T) {
	var gateway gateway.IGateway = GetMockGateway()
	b := bytes.NewBufferString("")

	command := get.NewCmdGet(gateway, "1.0")

	command.SetArgs([]string{})
	command.SetOut(b)
	command.SetErr(b)
	command.Execute()

	out, err := ioutil.ReadAll(b)
	if err != nil {
		t.Fatal(err)
	}

	assert.Equal(t, string(out), "Error: accepts 1 arg(s), received 0\nUsage:\n  get <address> [flags]\n\nAliases:\n  get, fetch, g\n\nFlags:\n  -a, --api string         set discovery api host on the fly only for this request\n  -b, --block-height int   gets an account at the given block height (default -1)\n  -f, --filter string      filter output to show only provided field (address, balance, code, keys)\n  -h, --help               help for get\n  -j, --json               show output format as JSON\n\n")
}
