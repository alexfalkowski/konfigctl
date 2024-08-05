[![CircleCI](https://circleci.com/gh/alexfalkowski/konfigctl.svg?style=svg)](https://circleci.com/gh/alexfalkowski/konfigctl)
[![codecov](https://codecov.io/gh/alexfalkowski/konfigctl/graph/badge.svg?token=QSRFU8VNST)](https://codecov.io/gh/alexfalkowski/konfigctl)
[![Go Report Card](https://goreportcard.com/badge/github.com/alexfalkowski/konfigctl)](https://goreportcard.com/report/github.com/alexfalkowski/konfigctl)
[![Go Reference](https://pkg.go.dev/badge/github.com/alexfalkowski/konfigctl.svg)](https://pkg.go.dev/github.com/alexfalkowski/konfigctl)
[![Stability: Active](https://masterminds.github.io/stability/active.svg)](https://masterminds.github.io/stability/active.html)

# Konfig Control

A tool to control the [konfig](https://github.com/alexfalkowski/konfig) daemon.

## Background

Have a look at the konfig [background](https://github.com/alexfalkowski/konfig?tab=readme-ov-file#background).

### Why a client?

We want to separate the daemon from the control. This is a similar design to other systems, such as [kubernetes](https://kubernetes.io/).

## Client

The client contains multiple commands. They all share the way we connect to the service, this is configured as:

```yaml
client:
  address: localhost:12000
  user_agent: "Konfig-client/1.0 gRPC/1.0"
  retry:
    attempts: 3
    backoff: 100ms
    timeout: 10s
  timeout: 5s
```

### Config

The client can download a configuration.

```bash
❯ ./konfigctl config --help
Get Config.

Usage:
  konfigctl config [flags]

Flags:
  -h, --help            help for config
  -o, --output string   output config location (format kind:location) (default "env:KONFIG_APP_CONFIG_FILE")

Global Flags:
  -i, --input string   input config location (format kind:location) (default "env:KONFIG_CONFIG_FILE")
```

This can be configured as following:

```yaml
client:
  config:
    application: test
    version: v1.11.0
    environment: staging
    continent: '*'
    country: '*'
    command: server
    kind: yaml
    mode: 0o600
```

### Secrets

The client can write secrets to a specified path.

```bash
❯ ./konfigctl config --help
Get Config.

Usage:
  konfigctl config [flags]

Flags:
  -h, --help            help for config
  -o, --output string   output config location (format kind:location) (default "env:KONFIG_APP_CONFIG_FILE")

Global Flags:
  -i, --input string   input config location (format kind:location) (default "env:KONFIG_CONFIG_FILE")
```

This can be configured as following:

```yaml
client:
  secrets:
    files:
      vault.secret: vault:/secret/data/transport/http/user_agent
      ssm.secret: ssm:/secret/data/transport/http/user_agent
    path: reports
    mode: 0o600
```

### Dependencies

![Dependencies](./assets/client.png)

## Design

Please take a look at the [template](https://github.com/alexfalkowski/go-client-template) this is derived from.

## Development

If you would like to contribute, here is how you can get started.

### Structure

The project follows the structure in [golang-standards/project-layout](https://github.com/golang-standards/project-layout).

### Dependencies

Please make sure that you have the following installed:
- [Ruby](.ruby-version)
- [Golang](go.mod)

### Style

This project favours the [Uber Go Style Guide](https://github.com/uber-go/guide/blob/master/style.md)

### Setup

Check out [CI](.circleci/config.yml).


### Changes

To see what has changed, please have a look at `CHANGELOG.md`
