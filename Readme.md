# Flow CLI
Flow CLI is a command line tool to interact with flow emulator and query data.

## Usage
  ```bash
  flow <command> <subcommand> [flags]
  ```

### Core Commands
```bash
  account:     Query account data from Flow
  config:      Configure CLI properties (such as API of emulator)
  help:        Help about any command
  version:     Show version information
```

### Examples
  ```bash
  $ flow account get "01cf0e2f2f715450" # get all account information
  $ flow account get "01cf0e2f2f715450" --filter address # get only account address
  $ flow account get "01cf0e2f2f715450" --json # get all account information in JSON
  $ flow config
  ```
  
## Demo
[![asciicast](https://asciinema.org/a/badcgVvoh6BjHlOGKHqK2jgaN.svg)](https://asciinema.org/a/badcgVvoh6BjHlOGKHqK2jgaN)

# Instalation

Download a binary for your OS at the [releases page](https://github.com/sideninja/flow-cli).