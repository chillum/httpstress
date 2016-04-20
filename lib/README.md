## httpstress library

httpstress is a Go library for HTTP stress testing.
It launches one goroutine per concurrent connection

It follows HTTP redirects.
Non-200 HTTP return codes are considered as errors

### Installation
* Install [Git](http://git-scm.com/download)
* Install [Go runtime](http://golang.org/doc/install).
  Go 1.3.3 on amd64 is recommended because of performance issues
* Set [`GOPATH`](http://golang.org/doc/code.html#GOPATH)
* `go get github.com/chillum/httpstress/lib`

### Docs
* [godoc.org](https://godoc.org/github.com/chillum/httpstress/lib)
* `godoc github.com/chillum/httpstress/lib`
