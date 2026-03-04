# Logging for Go

- Reads LOGLEVEL environment variable automatically to filter logs
- LOGLEVEL can be updated at runtime if 
- Default logger has log level as prefix and the log level uses ANSI colors
- NewLogger allows creating custom loggers with custom prefixes and suffixes as well as optional timestamp prefixes.

## Development

**Requires at least Go 1.26.0**

1. Install dependencies
    ```sh
    go mod download
    ```

1. Generate updated String method for LogLevel if it is changed
    ```sh
    go Generate ./...
    ```

1. Run tests
    ```sh
    go test ./...
    ```
