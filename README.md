# TopData Package Service

## About
- This is cli command written in Go that handles "Topdata Packages"
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
5. Compile and run the program (fast!):
```bash
go run .
```

## Install binary on your machine
Assuming you have Go installed on your machine, you can install the binary with the following command:
```bash
go install github.com/topdata-software-gmbh/topdata-package-util@latest
```
this compiles the program and installs it in your `$GOPATH/bin` directory.


## CLI Commands
Run the program for exploring the commands:
```bash
go run main.go --help
```

## Quick Start
```bash
# print a list of all packages as a table
main pkg list --all
```


## Command Line Options
- `--webserver-config-file`
  - Set the path to the config file. Default is `webserver-config.yaml`.
- `--portfolio-file`
  - Set the path to the config file where the package list is defined. Default is `portfolio.yaml`. 



## Webserver
The program has a built-in webserver for serving a JSON API, start it with:
```bash
go run main.go webserver
```

### API Endpoints

- `http://localhost:8080/` - Welcome message
- `http://localhost:8080/ping` - Pong
- `http://localhost:8080/repositories` - Returns a list of repositories
- `http://localhost:8080/repository-details/:name` - Returns details of a repository





## TODO
- fix and refactor the webservice API
- stats, see for example:
    - https://api.friendsofshopware.com/v2/packagist/packages
    - https://api.friendsofshopware.com/v2/shopware/sales
- pkg details: show sw6 store backend url: https://account.shopware.com/producer/plugins/123456

## Security issues
If you think that you have found a security issue, please contact security@topdata.de


## CHANGELOG:
2024-05-19: project name changed topdata-package-service -> topdata-package-util
2024-11-02: removed the package building stuff, as this is now handled by the topdata-package-release-builder

## Documentation

For more detailed information, please refer to the [documentation](./docs/index.md).