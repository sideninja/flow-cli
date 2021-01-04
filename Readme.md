![Flow](./flowcli.svg "Flow")

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

## Guidlines
Read guidlines for this CLI design: [OUR GUIDELINES](https://github.com/sideninja/flow-cli/blob/master/GUIDELINES.md)

## Demo
[![asciicast](https://asciinema.org/a/badcgVvoh6BjHlOGKHqK2jgaN.svg)](https://asciinema.org/a/badcgVvoh6BjHlOGKHqK2jgaN)

# Instalation

Download a binary for your OS at the [releases page](https://github.com/sideninja/flow-cli/releases/tag/0.1.0).

## Docker
You can use Docker to build and run the image. Be careful docker container can access the network of Flow Emulator. Create a shared network or use host if needed.
```bash
docker build -t flow .
docker run -it flow
```