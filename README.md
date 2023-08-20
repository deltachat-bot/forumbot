# Forum

![Latest release](https://img.shields.io/github/v/tag/deltachat-bot/forumbot?label=release)
[![CI](https://github.com/deltachat-bot/forumbot/actions/workflows/ci.yml/badge.svg)](https://github.com/deltachat-bot/forumbot/actions/workflows/ci.yml)
![Coverage](https://img.shields.io/badge/Coverage-20.0%25-red)
[![Go Report Card](https://goreportcard.com/badge/github.com/deltachat-bot/forumbot)](https://goreportcard.com/report/github.com/deltachat-bot/forumbot)

## Install

Binary releases can be found at: https://github.com/deltachat-bot/forumbot/releases

To install from source:

```sh
go install github.com/deltachat-bot/forumbot@latest
```

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
