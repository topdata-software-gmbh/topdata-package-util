# TopData Package Service

This is a microservice written in Go that handles git repositories from GitHub and Bitbucket with multiple release branches.

## Installation
```bash
# fetch dependencies 
go mod tidy
```

## Running the Service

1. Ensure you have Go installed on your machine.
2. Clone this repository.
3. Navigate to the project directory.
4. Run the server:
```bash
go run cmd/server/main.go
```

## Command Line Options

- `--port`
  - Set the port to run the server on. Default is `8080`.
  - example: `go run cmd/server/main.go --port=8081`
- `--config`
  - Set the path to the config file. Default is `config.json5`.
  - example: `go run cmd/server/main.go --config=path/to/config.json5`

  
## API Endpoints

- `http://localhost:8080/` - Welcome message
- `http://localhost:8080/ping` - Pong
- `http://localhost:8080/repositories` - Returns a list of repositories
- `http://localhost:8080/repository-details/:name` - Returns details of a repository


## TODO
- rename repository to package
- /get-release-branches/{packageMachineName}
- /get-releases/{packageMachineName}/{releaseBranchName}
