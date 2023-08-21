# Forum

![Latest release](https://img.shields.io/github/v/tag/deltachat-bot/forumbot?label=release)
[![CI](https://github.com/deltachat-bot/forumbot/actions/workflows/ci.yml/badge.svg)](https://github.com/deltachat-bot/forumbot/actions/workflows/ci.yml)
![Coverage](https://img.shields.io/badge/Coverage-20.0%25-red)
[![Go Report Card](https://goreportcard.com/badge/github.com/deltachat-bot/forumbot)](https://goreportcard.com/report/github.com/deltachat-bot/forumbot)

## Install

Binary releases can be found at: https://github.com/deltachat-bot/forumbot/releases

### Installing deltachat-rpc-server

This program depends on a standalone Delta Chat RPC server `deltachat-rpc-server` program that must be
available in your `PATH`. For installation instructions check:
https://github.com/deltachat/deltachat-core-rust/tree/master/deltachat-rpc-server

## Running the bot

Configure the bot:

```sh
forumbot init bot@example.com PASSWORD
```

Start listening to incoming messages:

```sh
forumbot serve
```

Run `forumbot --help` to see all available options.

## Contributing

For development instructions of the frontend check `frontend/README.md`

### Automated testing

You need to have a local fake email server running. The easiest way to do that is with Docker:

```
$ docker pull ghcr.io/deltachat/mail-server-tester:release
$ docker run -it --rm -p 3025:25 -p 3110:110 -p 3143:143 -p 3465:465 -p 3993:993 ghcr.io/deltachat/mail-server-tester
```

To run the automated tests run the script `./scripts/run_tests.sh`

### Manual testing

After building the frontend, to build and run the bot, execute in the project root folder:

```
go run ./...
```

### Building from source

After building the frontend, to build the bot, execute in the project root folder:

```
go build -o dist/ ./...
```

### Releasing

To automatically build and create a new GitHub release:

```
git tag v1.0.1
git push origin v1.0.1
```
