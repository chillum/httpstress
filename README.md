## httpstress-go

CLI utility for stress testing of HTTP servers with many concurrent connections.

Returns `0` if no errors, `1` if some requests failed, `2` on kill, `3` in case of invalid options
and `4` if it encounters a `setrlimit(2)`/`getrlimit(2)` error.

Prints elapsed time and error count for each URL to stdout (if any; does not count successful attempts).
Usage and runtime errors go to stderr.

#### Install: [source code](https://github.com/chillum/httpstress-go/wiki/Building-from-source) or [binary release](https://github.com/chillum/httpstress-go/wiki/Installing-from-binaries) (recommended) 

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
* Output is YAML-formatted. Example:
```yaml
Errors:
  - Location: http://localhost
    Count:    334
  - Location: http://127.0.0.1
    Count:    333
Elapsed time: 6.791903888s
```
