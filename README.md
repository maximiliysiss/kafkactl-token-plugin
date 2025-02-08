# kafkactl-token-plugin

`kafkactl-token-plugin` is an authentication plugin for [`kafkactl`](https://github.com/deviceinsight/kafkactl) that
provides simple token-based authentication (Bearer). The plugin is implemented in Go and can be used
to authenticate Kafka requests.

## Installation

To install the plugin, use the following command:

```sh
go install github.com/maximiliysiss/kafkactl-token-plugin@latest
```

Ensure that the binary is available in your system's `PATH` after installation.

## Usage

This plugin integrates with `kafkactl` to provide token-based authentication. Once installed, you can configure
`kafkactl` to use the plugin by specifying it in the `kafkactl` configuration.

Example configuration in `~/.kafkactl/config.yaml`:

```yaml
sasl:
  enabled: true
  mechanism: oauth
  tokenprovider:
    plugin: token
    options:
      token: {token}
```

## Development

To build the plugin locally:

```sh
git clone https://github.com/maximiliysiss/kafkactl-token-plugin.git
cd kafkactl-token-plugin
go build -o kafkactl-token-plugin .
```

To run tests:

```sh
go test ./...
```

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

