## httpstress-go [![Build Status](https://travis-ci.org/chillum/httpstress-go.svg?branch=master)](https://travis-ci.org/chillum/httpstress-go)

httpstress-go is a CLI utility for stress testing of HTTP servers with many concurrent connections.

Returns 0 if no errors, 1 if some failed (see stdout), 2 on kill and 3 in case of invalid options.

Prints error count for each URL to stdout (does not count successful attempts).
Errors and debugging information go to stderr.

### Installing from binaries
Extract the appropriate archive and launch `httpstress-go` with desired options

* Windows (compiled on Windows 7 SP1)
  * [64-bit Windows](../../releases/download/v2.0.0.1/win64.zip) (recommended)
  * [32-bit Windows](../../releases/download/v2.0.0.1/win32.zip)
* [Mac OS X](../../releases/download/v2.0.0.1/mac.zip) (compiled on a 10.8 system)
* Linux (compiled on CentOS 6.5)
  * [64-bit Linux](../../releases/download/v2.0.0.1/linux64.tgz) (recommended)
  * [32-bit Linux](../../releases/download/v2.0.0.1/linux32.tgz)

### Installing from source
* Supported platforms: Unix (Mac OS X, Linux, FreeBSD) and Windows
* Install [Git](http://git-scm.com/download)
* Install [Go runtime](http://golang.org/doc/install).
  64-bit Go 1.3 or higher is recommended because of performance issues
* Set [`GOPATH`](http://golang.org/doc/code.html#GOPATH)
* `go get github.com/chillum/httpstress-go`
* Ready to use: launch `httpstress-go` with desired options

### Notes
* This ulility takes care of `ulimit -n` on Unix systems: sets it to
  the value of `-c` option plus 6, if the current limit is smaller.
* Error output is YAML-formatted. Example:
```yaml
errors:
  - location: http://localhost
    count:    334
  - location: http://127.0.0.1
    count:    333
```
* `httpstress-go` is
  [a static-linked binary](http://golang.org/doc/faq#Why_is_my_trivial_program_such_a_large_binary),
  it's possible to deploy it just by copying `$GOPATH/bin/httpstress-go`
  (`%GOPATH%\bin\httpstress-go.exe` on Windows),
  compiled on matching system and architecture

### Environment
`GOMAXPROCS` – Go threads number (defaults to CPU count + 1)

### Usage
`httpstress-go [options] <URL list>`

### Options
* `URL list` – URLs to fetch (required)
* `-c NUM` – concurrent connections number (defaults to 1)
* `-n NUM` – total connections number (optional)
* `-v` – print version to stdout and exit

### Example usage
`httpstress-go -c 1000 http://localhost http://google.com`
