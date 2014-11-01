/* CLI utility for stress testing of HTTP servers with many concurrent connections

Usage:
 httpstress-go [options] <URL list>

Options:
 * `URL list` – URLs to fetch (required)
 * `-c NUM` – concurrent connections number (defaults to 1)
 * `-n NUM` – total connections number (optional)
 * `-v` – print version to stdout and exit

Example:
 httpstress-go -c 1000 http://localhost https://google.com

Returns 0 if no errors, 1 if some failed (see stdout), 2 on kill and 3 in case of invalid options.

Prints error count for each URL to stdout (does not count successful attempts).
Errors and debugging information go to stderr.

Error output is YAML-formatted. Example:
 errors:
   - location: http://localhost
     count:    334
   - location: http://127.0.0.1
     count:    333

Please note that this utility uses GOMAXPROCS environment variable if it's present.
If not, this defaults to CPU count + 1. */
package main

/* Copyright 2014 Vasily Korytov

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

// Application version
const Version = "2.1"

func main() {
	var conn, max int
	flag.IntVar(&conn, "c", 1, "concurrent connections count")
	flag.IntVar(&max, "n", 0, "total connections (optional)")
	version := flag.Bool("v", false, "print version to stdout and exit")
	flag.Usage = func() {
		fmt.Fprintln(os.Stderr, "Usage:", os.Args[0], "[options] <URL list>")
		fmt.Fprintln(os.Stderr, "  <URL list>: URLs to fetch (required)")
		flag.PrintDefaults()
		fmt.Fprintln(os.Stderr, "Docs:\n  https://github.com/chillum/httpstress-go")
		fmt.Fprintln(os.Stderr, "Example:\n  httpstress-go -c 1000 http://localhost https://google.com")
		os.Exit(3)
	}
	flag.Parse()

	if *version {
		fmt.Println("httpstress-go", Version)
		fmt.Println("httpstress", httpstress.Version)
		fmt.Println(runtime.Version(), runtime.GOOS, runtime.GOARCH)
		os.Exit(0)
	}

	urls := flag.Args()
	if len(urls) < 1 {
		flag.Usage()
	}

	if os.Getenv("GOMAXPROCS") == "" {
		runtime.GOMAXPROCS(runtime.NumCPU() + 1)
	}

	setlimits(&conn) // Platform-specific code: see unix.go and windows.go for details.

	out, err := httpstress.Test(conn, max, urls)
	if err != nil {
		fmt.Fprintln(os.Stderr, "ERROR:", err)
		flag.Usage()
	}

	if len(out) > 0 {
		fmt.Println("errors:")
		for url, num := range out {
			fmt.Println("  - location: ", url, "\n    count:    ", num)
		}
		os.Exit(1)
	} else {
		fmt.Fprintln(os.Stderr, "Test finished. No failed requests.")
	}
}
