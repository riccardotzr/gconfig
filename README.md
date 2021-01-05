<div align="center">

# gconfig

![Build Status](https://github.com/riccardoAtzori91/gconfig/workflows/build/badge.svg)

</div>

GConfig is a go library to handle configuration file and environment variables.

It uses [viper](https://github.com/spf13/viper) to handle env vars and configuration file.

## Install

```ssh
go get -u github.com/riccardoatzori/gconfig
```

## Usage

### Read Environment variables

```go
type EnvironmentVariables struct {
    LogLevel string
    HTTPPort string
}

var envVariablesConfig = []EnvConfig{
    {
        Key:      "LOG_LEVEL",
        Variable: "LogLevel",
    },
    {
        Key:      "HTTP_PORT",
        Variable: "HTTPPort",
    },
}

var env TestEnvironmentVariables
err := GetEnvVariables(envVariablesConfig, &env)
```

## License

This project is licensed under the Apache License 2.0 - see the [LICENSE.md](LICENSE.md)
file for details