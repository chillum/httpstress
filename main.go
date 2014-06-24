/* httpstress-go is a CLI utility for stress testing of HTTP servers with many concurrent connections.

Usage: httpstress-go -c {concurrent} -n {total} {URL list}
e.g. httpstress-go -c 1000 -n 2000 http://localhost http://google.com

{concurrent} defaults to 1, {total} is optional.

Returns 0 if no errors, 1 if some failed (see stdout), 2 on kill and 3 in case of invalid options.

Prints error count for each URL to stdout (does not count successful attempts).
Errors and debugging information go to stderr.

Please note that this utility uses GOMAXPROCS environment variable if it's present.
If not, this defaults to CPU count + 1.

Be sure to set `ulimit -n` on Unix systems (Mac OS X, Linux, FreeBSD).*/
package main

/* Copyright 2014 Chai Chillum

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License. */

import (
	"flag"
	"fmt"
	"github.com/chillum/httpstress"
	"os"
	"runtime"
)

func main() {
	var conn, max int
	flag.IntVar(&conn, "c", 1, "concurrent connections count")
	flag.IntVar(&max, "n", 0, "total connections (optional)")
	flag.Usage = func() {
		// TODO: stderr
		fmt.Fprintln(os.Stderr, "Usage:", os.Args[0], "[options] <http://url1> [http://url2] ... [http://urlN]")
		flag.PrintDefaults()
		fmt.Fprintln(os.Stderr, "Docs:\n  https://godoc.org/github.com/chillum/httpstress-go")
		fmt.Fprintln(os.Stderr, "  godoc github.com/chillum/httpstress-go")
		os.Exit(3)
	}
	flag.Parse()
	urls := flag.Args()
	if len(urls) < 1 {
		flag.Usage()
	}

	if os.Getenv("GOMAXPROCS") == "" {
		runtime.GOMAXPROCS(runtime.NumCPU() + 1)
	}

	out, err := httpstress.Test(conn, max, urls)
	if err != nil {
		fmt.Fprintln(os.Stderr, "ERROR:", err)
		os.Exit(3)
	}
	if len(out) > 0 {
		fmt.Fprintln(os.Stderr, "Test finished. Failed requests:")
		for url, num := range out {
			fmt.Print(" ", url, ": ", num, "\n")
		}
		os.Exit(1)
	} else {
		fmt.Fprintln(os.Stderr, "Test finished. No failed requests.")
	}
}
