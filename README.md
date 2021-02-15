<p align="center">
<img src="logo.svg" alt="Alert Logo">
</p>

# alert
[![GoDoc][doc-img]][doc] [![Go Report Card][report-card-img]][report-card] [![GolangCI Lint][golangci-lint-img]][golangci-lint]


A Simple cross-platform alert library for go.

## Installation


`go get github.com/fcjr/alert`


## Usage

```go
import "github.com/fcjr/alert"

alert.Info("Example Message", "Example Message Text")
```

## Progress

* [X] Basic mac
* [X] Basic windows
* [ ] Basic linux
* [ ] Advanced mac (Input, Scroll UI, etc)
* [ ] Advanced windows
* [ ] Advanced linux

[doc-img]: https://img.shields.io/static/v1?label=godoc&message=reference&color=blue
[doc]: https://pkg.go.dev/github.com/fcjr/alert?tab=doc
[report-card-img]: https://goreportcard.com/badge/github.com/fcjr/alert
[report-card]: https://goreportcard.com/report/github.com/fcjr/alert
[golangci-lint-img]: https://github.com/fcjr/alert/workflows/golangci-lint/badge.svg
[golangci-lint]: https://github.com/fcjr/alert/actions?query=workflow%3Agolangci-lint