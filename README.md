# TopData Package Service

This is a microservice written in Go that handles git repositories from GitHub and Bitbucket with multiple release branches.

## Installation
```bash
go mod tidy
```

## Prod URL
http://packages.api.topinfra.de

## Running the Service

1. Ensure you have Go installed on your machine.
2. Clone this repository.
3. Navigate to the project directory.
4. Run `go run .`.

## API Endpoints

- `http://localhost:8080/` - Welcome message
- `http://localhost:8080/repositories` - Returns a list of repositories

## Command Line Options

- `--port`
    - Set the port to run the server on. Default is `8080`.
    - example: `go run . --port=8081`

