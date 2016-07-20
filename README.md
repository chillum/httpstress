## httpstress

CLI utility for stress testing of HTTP servers with many concurrent connections.

Prints elapsed time and error count for each URL to stdout (if any).  
Produces JSON-formatted output like:

```json
{
  "errors": {
    "http://localhost": 500,
    "https://192.168.1.1": 3
  },
  "seconds": 12.8
}
```

#### Install: [source code](https://github.com/chillum/httpstress/wiki/Installing-from-source) or [binary release](https://github.com/chillum/httpstress/wiki/Installing-from-binaries)

#### Use: check out [the user manual](https://github.com/chillum/httpstress/wiki)
