/*
CLI utility for stress testing of HTTP servers with many concurrent connections

Usage:
 httpstress [options] <URL list>

Options:
 * -c <int>   concurrent connections number (defaults to 1)
 * -n <int>   total connections number (defaults to URL count)
 * --version  print version to stdout and exit

Example:
 httpstress -c 1000 http://localhost https://192.168.1.1

Returns 0 if no errors, 1 if some requests failed, 2 on kill and 3 in case of invalid options.

Prints elapsed time and error count for each URL to stdout (if any; does not count successful attempts).
Usage and runtime errors go to stderr.

Output is JSON-formatted. Example:
  {
    "errors": {
      "http://localhost": 500,
      "https://192.168.1.1": 3
    },
    "seconds": 12.8
  }

It follows HTTP redirects. Non-200 HTTP return code is an error.

Be sure to set `ulimit -n` on Unix systems high enough.
*/
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"runtime"
	"time"

	httpstress "github.com/chillum/httpstress/lib"
)

// Version is the application version
const Version = "6.5.2"

type results struct {
	Errors  interface{} `json:"errors"`
	Seconds *float32    `json:"seconds"`
}

type ver struct {
	App  string `json:"httpstress"`
	Go   string `json:"runtime"`
	Os   string `json:"os"`
	Arch string `json:"arch"`
}

func main() {
	var conn, max int
	var final results
	flag.IntVar(&conn, "c", 1, "concurrent connections")
	flag.IntVar(&max, "n", 0, "total connections (defaults to URL count)")
	version := flag.Bool("version", false, "print version to stdout and exit")
	flag.Usage = func() {
		fmt.Fprintln(os.Stderr, "Usage: httpstress [options] <URL list>\nOptions:")
		flag.PrintDefaults()
		fmt.Fprintln(os.Stderr, "Example: httpstress -c 1000 http://localhost https://192.168.1.1")
		fmt.Fprintln(os.Stderr, "Docs:    https://github.com/chillum/httpstress/wiki")
		os.Exit(3)
	}
	flag.Parse()

	if *version {
		var ver ver
		ver.App = Version
		ver.Go = runtime.Version()
		ver.Os = runtime.GOOS
		ver.Arch = runtime.GOARCH
		json, _ := json.Marshal(&ver)
		fmt.Println(string(json))
		os.Exit(0)
	}

	urls := flag.Args()
	if len(urls) < 1 {
		flag.Usage()
	}

	start := time.Now()

	errors, err := httpstress.Test(conn, max, urls)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Error:", err)
		flag.Usage()
	}

	elapsed := float32(int64(time.Since(start).Seconds()*10)) / 10

	if len(errors) > 0 {
		defer os.Exit(1)
	}

	final.Errors = &errors
	final.Seconds = &elapsed

	json, _ := json.MarshalIndent(&final, "", "  ")
	fmt.Println(string(json))
}
