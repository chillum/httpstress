## httpstress-go
httpstress-go is a CLI interface for
[httpstress](https://github.com/chillum/httpstress.git) library.

Use it for stress testing of HTTP servers with many concurrent connections.

Returns 0 if no errors, 1 if some errors (see stdout) and 2 in case of invalid options.

### Installation
* Install [Go runtime](http://golang.org/doc/install)
* `go get github.com/chillum/httpstress-go`
* Ready to use: launch `httpstress-go` with desired options

### Options
* `-c NUM` -- concurrent connections number (defaults to 1)
* `-n NUM` -- total connections number (optional)
* `URL list` -- URLs to fetch

### Example usage
`httpstress-go -c 1000 -m 2000 http://localhost http://google.com`
