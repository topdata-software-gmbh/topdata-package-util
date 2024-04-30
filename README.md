# TopData Package Service

This is a microservice written in Go that handles git repositories from GitHub and Bitbucket with multiple release branches.


## Running the Service

1. Ensure you have Go installed on your machine.
2. Clone this repository.
3. Navigate to the project directory.
4. Run `go run main.go`.

## API Endpoints

- `http://localhost:8080/` - Welcome message
- `http://localhost:8080/repositories` - Returns a list of repositories
