## httpstress-go

CLI utility for stress testing of HTTP servers with many concurrent connections.  
Automatically sets `ulimit -n` on Unix systems.

Prints elapsed time and error count for each URL to stdout (if any).  
Produces YAML-formatted output like:

```yaml
Errors:
  - Location: http://localhost
    Count:    334
  - Location: https://127.0.0.1
    Count:    333
Elapsed time: 4.791903888s
```

#### Install: [source code](https://github.com/chillum/httpstress-go/wiki/Building-from-source) or [binary release](https://github.com/chillum/httpstress-go/wiki/Installing-from-binaries) (recommended) 

#### Use: check out [the user manual](https://github.com/chillum/httpstress-go/wiki#httpstress-go)
