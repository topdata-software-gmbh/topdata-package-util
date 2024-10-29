# Commands

This document provides an overview of all available commands in the application.

## Git Repository Management

- `pkg fetch` - Fetches updates from remote repositories
- `pkg clone` - Clones repositories that haven't been cloned yet
- `pkg clean` - Removes all local git repositories

## Package Management

- `pkg list` - Lists all configured packages and their branches
- `pkg details [packageName]` - Shows detailed information about a specific package
- `pkg build-release-zip [packageName] [releaseBranchName]` - Builds a release zip for uploading to the shopware6 plugin store
- `pkg find-branch` - Finds the branch with the highest plugin version for a given Shopware version
- `pkg test-git` - Testing git cli wrapper

## Branch Management

- `pkg show-git-branch-details [packageName] [branchName]` - Shows details of a single branch of a repository

## Cache Management

- `cache clear` - Clears the cache

## Local Git Repository Management

- `localgit compare-branches [branch1] [branch2]` - Compares two branches and shows the differences in a table

## Information Commands

- `version` - Displays the current version of the tool
- `ping` - Just a test command

## Web Interface

- `webserver` - Starts the web interface server

## Configuration

- `config` - Shows the current configuration
- `init` - Initializes a new configuration file

For detailed information about each command, use the `--help` flag with any command.
