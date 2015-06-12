## httpstress-go [![Build Status](https://travis-ci.org/chillum/httpstress-go.svg?branch=master)](https://travis-ci.org/chillum/httpstress-go)

CLI utility for stress testing of HTTP servers with many concurrent connections.

Returns `0` if no errors, `1` if some failed (see stdout), `2` on kill, `3` in case of invalid options
and `4` if it encounters a `setrlimit(2)`/`getrlimit(2)` error.

Prints error count for each URL to stdout (does not count successful attempts).
Errors and debugging information go to stderr.

### [Installing from source](https://github.com/chillum/httpstress-go/wiki/Building-from-source)

### Installing from binaries
Extract the appropriate archive and launch `httpstress-go` with desired options

* Windows
  * [64-bit Windows](https://github.com/chillum/httpstress-go/releases/download/v3.0/win64.zip) (recommended)
  * [32-bit Windows](https://github.com/chillum/httpstress-go/releases/download/v3.0/win32.zip)
* [Mac OS X](https://github.com/chillum/httpstress-go/releases/download/v3.0/mac.zip) (10.7 or greater)
* Linux
  * [x86-64 Linux](https://github.com/chillum/httpstress-go/releases/download/v3.0/linux_amd64.zip) (recommended)
  * [i386 Linux](https://github.com/chillum/httpstress-go/releases/download/v3.0/linux_386.zip)

### Environment
`GOMAXPROCS` – Go threads number (defaults to CPU count + 1)

### Usage
`httpstress-go <URL list> [options]`

### Options
* `URL list` – URLs to fetch (required)
* `-c NUM` – concurrent connections number (defaults to 1)
* `-n NUM` – total connections number (optional)
* `-v` – print version to stdout and exit

### Example usage
`httpstress-go http://localhost https://google.com -c 1000`

### Notes
* This ulility takes care of `ulimit -n` on Unix systems: sets it to
  the value of `-c` option plus 6, if the current limit is smaller.
* Error output is YAML-formatted. Example:
```yaml
Errors:
  - Location: http://localhost
    Count:    334
  - Location: http://127.0.0.1
    Count:    333
```
