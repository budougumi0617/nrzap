# nrzap

[![Go Reference](https://pkg.go.dev/badge/github.com/budougumi0617/nrzap.svg)](https://pkg.go.dev/github.com/budougumi0617/nrzap)
[![MIT License](http://img.shields.io/badge/license-MIT-blue.svg?style=flat-square)](LICENSE)
[![test](https://github.com/budougumi0617/nrzap/workflows/test/badge.svg)](https://github.com/budougumi0617/nrzap/actions?query=workflow%3Atest)
[![reviewdog](https://github.com/budougumi0617/nrzap/workflows/reviewdog/badge.svg)](https://github.com/budougumi0617/nrzap/actions?query=workflow%3Areviewdog)

## Description
New Relic Logs in context for the Go agent connects your logs and APM data in New Relic. but, official New Relic agent povides a plugin only for logurus.

- https://github.com/newrelic/go-agent/tree/master/v3/integrations/logcontext

nrzap is helper for go.uber.org/zap logger. nrzap helps to connect logs and APM data in New Relic.
## Prerequisites

- Enable distributed tracing for your Go applications
    - https://docs.newrelic.com/jp/docs/agents/go-agent/features/enable-distributed-tracing-your-go-applications/
- Enable log management in New Relic
    - https://docs.newrelic.com/docs/logs/enable-log-management-new-relic/enable-log-monitoring-new-relic/enable-log-management-new-relic/

## Usage
`GetNrMetadataFields` funcion generates zap typed fields for connect logs and APM in New Relic.

```go
func ExampleHandler(w http.ResponseWriter, r *http.Request) {
	logger, _ := zap.NewProduction()
    defer logger.Sync()
    nrfs := nrzap.GetNrMetadataFields(ctx)
    logger.Info("failed to fetch URL",
        // Structured context as strongly typed Field values.
        zap.String("url", url),
        zap.Int("attempt", 3),
        zap.Duration("backoff", time.Second),
        nrfs...,
    )
}
```
## Installation

```bash
$ go get -u github.com/budougumi0617/nrzap
```

## License

[MIT](./LICENSE)

## Author
Yocihiro Shimizu(@budougumi0617)