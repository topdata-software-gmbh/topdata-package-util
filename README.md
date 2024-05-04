# TopData Package Service

## About
- This is cli command written in Go that handles "Topdata Packages"
- aka "Topdata Release Manager"
- a Topdata Package is currently a "Shopware 6 Plugin", more to come.
- git repositories with branches are used for release management 
- it is single CLI program with multiple commands and subcommands
- It has a webserver with JSON endpoints to be used for generating documentation pages (using mkdocs) and maybe for other services later

## Compile and run the program

1. Ensure you have Go installed on your machine.
2. Clone this repository.
3. Navigate to the project directory.
4. Install the dependencies:
```bash
go mod tidy
```
5. Run the program:
```bash
go run .
```

## Command Line Options

- `--packages-portfolio-file`
  - Set the path to the config file. Default is `webserver-config.json5`.
- `--packages-portfolio-file`
  - Set the path to the config file where the package list is defined. Default is `packages-portfolio.json5`. 

  
## API Endpoints

- `http://localhost:8080/` - Welcome message
- `http://localhost:8080/ping` - Pong
- `http://localhost:8080/repositories` - Returns a list of repositories
- `http://localhost:8080/repository-details/:name` - Returns details of a repository


## TODO
- rename repository to package
- /get-release-branches/{packageMachineName}
- /get-releases/{packageMachineName}/{releaseBranchName}
