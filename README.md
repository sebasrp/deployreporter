# deploy-reporter

Tool to retrieve and visualize deployments, through different means: a cli, a web server and a web client.
Right now it consumes grafana annotations, but can be extended to display other events relatively easily.

## web server

Uses the [Gin](https://github.com/gin-gonic/gin) framework. Serves right now one API endpoint `\deployments` to retrieve the deployments.

## cli

A utility `deployreporter` CLI is provided, that exposes the module through a simple interface. Use the search function of the cli to get details.

```shell
$ deployreporter -h
Tired of gathering deployment information from difference sources?
deployreporter collects and aggregate deployment data to quickly find what has changed in your service

Usage:
  deployreporter [command]

Available Commands:
  get         Retrieves deployments
  help        Help about any command

Flags:
      --config string       config file (default is $HOME/.deployreporter.yaml)
      --console             output results to console
      --csv                 output results to a csv file
      --grafanaKey string   API key used to retrieve data from grafana

      --from string         epoch datetime in milliseconds
      --to string           epoch datetime in milliseconds
  -v, --verbose             enables verbose output
  -h, --help                help for deployreporter

Use "deployreporter [command] --help" for more information about a command.
```

### Usage

Make sure you provide the required access keys to access the different tools.

Check the help page with `deployreporter --help` to see all available commands.

### Configuration file

Tired of manually selecting the different parameters? You can save those in a file and provide it with the `--config` flag - or just place it under `$HOME/.deployreporter.yaml` to be automatically picked up. The format and options supported are (order does not matter)

```yaml
grafana_access_key: <grafana access key>
console: true /false
csv: true / false
verbose: true / false
```

## Development

To run the latest:

```shell
cd deployreporter
go build ./... && go install ./...
deployreporter --help
```
