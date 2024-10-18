# deploy-reporter

Tool to retrieve and visualize deployments

## cli

A utility `deploy-reporter` CLI is provided, that exposes the module through a simple interface.

### Usage

Make sure you provide the required access keys to access the different tools.

Check the help page with `deploy-reporter --help` to see all available commands.

### Configuration file

Tired of manually selecting the different parameters? You can save those in a file and provide it with the `--config` flag - or just place it under `$HOME/..deploy-reporter.yaml` to be automatically picked up. The format and options supported are (order does not matter)

```yaml
grafana_access_key: <name of profile>
console: true /false
csv: true / false
verbose: true / false
```

## Development

To run the latest:

```shell
cd deploy-reporter
go build ./... && go install ./...
deploy-reporter --help
```
