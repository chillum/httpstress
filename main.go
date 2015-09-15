/*
CLI utility for stress testing of HTTP servers with many concurrent connections

Usage:
 httpstress <URL list> [options]

Options:
 * `URL list` – URLs to fetch (required)
 * `-c NUM` – concurrent connections number (defaults to 1)
 * `-n NUM` – total connections number (optional)
 * `-v` – print version to stdout and exit

Example:
 httpstress http://localhost https://google.com -c 1000

Returns 0 if no errors, 1 if some requests failed, 2 on kill, 3 in case of invalid options
and 4 if it encounters a setrlimit(2)/getrlimit(2) error.

Prints elapsed time and error count for each URL to stdout (if any; does not count successful attempts).
Usage and runtime errors go to stderr.

Output is YAML-formatted. Example:
 Errors:
   - Location: http://localhost
     Count:    334
   - Location: https://127.0.0.1
     Count:    333
 Elapsed time: 4.791903888s
*/
package main

import (
	"fmt"
	"github.com/chillum/httpstress/lib"
	flag "github.com/ogier/pflag"
	"os"
	"runtime"
	"time"
)

// Application version
const Version = "5"

func main() {
	var conn, max int
	flag.IntVarP(&conn, "c", "c", 1, "concurrent connections count")
	flag.IntVarP(&max, "n", "n", 0, "total connections (optional)")
	version := flag.BoolP("version", "v", false, "print version to stdout and exit")
	flag.Usage = func() {
		fmt.Fprintln(os.Stderr, "Usage:", os.Args[0], "<URL list> [options]")
		fmt.Fprintln(os.Stderr, "  <URL list>: URLs to fetch (required)")
		flag.PrintDefaults()
		fmt.Fprintln(os.Stderr, "Docs:\n  https://github.com/chillum/httpstress/wiki")
		fmt.Fprintln(os.Stderr, "Example:\n  httpstress http://localhost https://google.com -c 1000")
		os.Exit(3)
	}
	flag.Parse()

	if *version {
		fmt.Println("cli:", Version, "\nlib:", httpstress.Version,
			"\ngo: ", runtime.Version(),"\nos: ", runtime.GOOS, "\ncpu:", runtime.GOARCH)
		os.Exit(0)
	}

	urls := flag.Args()
	if len(urls) < 1 {
		flag.Usage()
	}

	if os.Getenv("GOMAXPROCS") == "" {
		runtime.GOMAXPROCS(runtime.NumCPU())
	}

	if !setlimits(&conn) { // Platform-specific code: see unix.go and windows.go for details.
		os.Exit(4)
	}

	start := time.Now()

	out, err := httpstress.Test(conn, max, urls)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		flag.Usage()
	}

	elapsed := time.Since(start)

	if len(out) > 0 {
		fmt.Println("Errors:")
		for url, num := range out {
			fmt.Println("  - Location:", url, "\n    Count:   ", num)
		}
		defer os.Exit(1)
	}
	fmt.Println("Elapsed time:", elapsed)
}
