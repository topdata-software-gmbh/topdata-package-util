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


## CLI Commands
Run the program for exploring the commands:
```bash
go run main.go --help
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



### Shopware Plugin Versioning Scheme
- [MAJOR].[MINOR].[PATCH]
- when a new release is created, the version number is increased by 1 in the following way:
  - MAJOR: increased when there are breaking changes
  - MINOR: increased when there are new features
  - PATCH: increased when there are bug fixes 

## Similar (?) Projects 
- https://github.com/pickware/scs-commander/ [javascript]
- https://github.com/shopwareLabs/plugin-info [php]
- https://github.com/shopwareLabs/sw-cli-tools [php]
- https://github.com/FriendsOfShopware/shopware-cli [golang]
- https://github.com/FriendsOfShopware/api.friendsofshopware.com [golang]
 


## TODO
- when creating a relaease zip, log it somewhere (release-log-path should be part of the config file)
- fix and refactor the webservice API
- make use of .sw-zip-blacklist when creating a release zip
    - example: https://github.com/shopware/SwagMigrationConnector/blob/master/.sw-zip-blacklist
- stats, see for example:
    - https://api.friendsofshopware.com/v2/packagist/packages
    - https://api.friendsofshopware.com/v2/shopware/sales


## Security issues
If you think that you have found a security issue, please contact security@topdata.de

