<div align="center" style="padding-top: 30px; padding-bottom: 30px">
    <h1><span style="color: aqua">L</span>ogger</h1>
</div>

[![License][License-Image]][License-URL]
[![Release][Release-Image]][Release-URL]

[License-URL]: https://www.apache.org/licenses/LICENSE-2.0
[License-Image]: https://img.shields.io/badge/License-Apache%202.0-blue.svg
[Release-URL]: https://github.com/RevittConsulting/logger/releases/latest
[Release-Image]: https://img.shields.io/github/v/release/RevittConsulting/logger?kill_cache=1

Logger is a go package that provides a simple and easy to use logging interface. It is designed to be used in a variety of applications and provides a simple way to log messages.

***

## Quick Start

```bash
go get github.com/RevittConsulting/logger
```
***

## Features

- [x] Simple and easy to use
- [x] Get observed logs for testing
- [x] Loki support

***

## Loki Configuration

Run this on app startup to enable loki logging.

```go
InitLoggerConfig(&Config{
    LokiAddress: "http://localhost:3100",
    Labels:      map[string]string{"app": "My App", "env": "local"},
    UsingLoki:   true,
})
```
***

## Log to file

Add these values to config

```go
loggerCfg := &logger.Config{
    LogToFile:   true,
    LogFilePath: "log.txt",
}
logger.InitLoggerConfig(loggerCfg)
```

***

## Contributing

Found a bug or want to suggest a feature? Feel free to create an issue or make a pull request.

***