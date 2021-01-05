<div align="center">

# GConfig

[![Build Status](https://github.com/riccardotzr/gconfig/workflows/build/badge.svg)](https://github.com/riccardotzr/gconfig/actions)
[![Go Report Card](https://goreportcard.com/badge/github.com/riccardotzr/gconfig)](https://goreportcard.com/report/github.com/riccardotzr/gconfig)
[![Go Reference](https://pkg.go.dev/badge/github.com/riccardotzr/gconfig.svg)](https://pkg.go.dev/github.com/riccardotzr/gconfig)

</div>

GConfig is a go library to handle configuration file and environment variables.

It uses [viper](https://github.com/spf13/viper) to handle env vars and configuration file.

## Install

```ssh
go get -u github.com/riccardotzr/gconfig
```

## Usage

### Read Environment variables

```go
type EnvironmentVariables struct {
    LogLevel string
    HTTPPort string
}

var envVariablesConfig = []gconfig.EnvConfig{
    {
        Key:      "LOG_LEVEL",
        Variable: "LogLevel",
    },
    {
        Key:      "HTTP_PORT",
        Variable: "HTTPPort",
    },
}

var env EnvironmentVariables

if err := gconfig.GetEnvVariables(envVariablesConfig, &env); err != nil {
    panic(err.Error())
}
```

### Read from file

```go
type ConfigurationFileVariables struct {
    LogLevel string
    HTTPPort string
}

var configuration ConfigurationFileVariables

if err := gconfig.GetConfigFromFile("my/path", "file", &configuration); err != nil {
    panic(err.Error())
}
    
```

## License

This project is licensed under the Apache License 2.0 - see the [LICENSE.md](LICENSE.md)
file for details