/* httpstress-go is a CLI interface for httpstress library. */
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
	. "fmt"
	. "github.com/chillum/httpstress"
	. "os"
)

func main() {
	var conn, max int
	flag.IntVar(&conn, "c", 1000, "concurrent connections count")
	flag.IntVar(&max, "m", 0, "total connections (optional)")
	flag.Parse()

	urls := flag.Args()
	if len(urls) < 1 {
		Println("Usage:", Args[0], "<http://url1> [http://url2] ... [http://urlN]")
		Exit(1)
	}

	out, err := Test(conn, max, urls)
	if err != nil {
		Println("ERROR:", err)
		Exit(1)
	}
	if len(out) > 0 {
		Println("Test finished. Failed requests:")
		for url, num := range out {
			Print(" ", url, ": ", num, "\n")
		}
	} else {
		Println("Test finished. No failed requests.")
	}
}
