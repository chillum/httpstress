## httpstress-go

[![Build Status](https://travis-ci.org/chillum/httpstress-go.svg?branch=master)](https://travis-ci.org/chillum/httpstress-go)

httpstress-go is a CLI interface for
[httpstress](https://github.com/chillum/httpstress.git) library.

Use it for stress testing of HTTP servers with many concurrent connections.

Returns 0 if no errors, 1 if some errors (see stdout) and 2 in case of invalid options.

Prints error count for each URL to stdout (does not count successful attempts).

### Installation
* Install [Go runtime](http://golang.org/doc/install).
  Go 1.3 or higher is recommended because of performance improvements
* Set [`GOPATH`](http://golang.org/doc/code.html#GOPATH)
* `go get github.com/chillum/httpstress-go`
* Ready to use: launch `httpstress-go` with desired options

### Environment
* `GOMAXPROCS` – Go threads number (defaults to CPU count + 1)

### Options
* `-c NUM` – concurrent connections number (defaults to 1)
* `-n NUM` – total connections number (optional)
* `URL list` – URLs to fetch

### Example usage
`httpstress-go -c 1000 -n 2000 http://localhost http://google.com`
