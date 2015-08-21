package main

import (
	"errors"
	"net/http"
	"strings"
)

/*
Test launches {conn} goroutines to fetch HTTP/HTTPS locations in {urls} list

If {max} is more than {conn}, more goroutines will spawn as other are finished,
resulting in {max} queries (but no more than {conn} in every moment).
Returns map: {url}, {fail count} or error (failed URL message). Example:

	urls := []string{"https://google.com", "http://localhost"}

	out, err := httpstress.Test(2, 4, urls) // Makes 4 HTTP requests total, 2 concurrent.

	if err != nil {
		// Invalid arguments.
	} else {
		if len(out) == 0 {
			// No failed requests.
		} else {
			// Process failed requests.
			for url, num := range out {
				log.Println(url, "failed", num, "times.")
			}
		}
	}
*/
func Test(conn int, max int, urls []string) (results map[string]int, err error) {
	for _, i := range urls {
		if !strings.HasPrefix(i, "http://") && !strings.HasPrefix(i, "https://") {
			err = errors.New("Not a HTTP/HTTPS URL: " + i)
			return
		}
	}

	// Ensure, every URL gets a request.
	if max < len(urls) {
		max = len(urls)
	}

	results = make(map[string]int)
	finished := make(chan string)
	total := len(urls) - 1
	trans := &http.Transport{MaxIdleConnsPerHost: conn} // Use persistent connections.
	client := &http.Client{Transport: trans}
	n := 0
	i := 0
	for ; i < conn; i++ { // Launch initial workers.
		go worker(&urls[n], finished, client)

		if n < total {
			n++
		} else {
			n = 0
		}
	}
	for ; i < max; i++ { // Launch more workers as initial finish.
		if url := <-finished; url != "" {
			results[url]++
		}

		go worker(&urls[n], finished, client)

		if n < total {
			n++
		} else {
			n = 0
		}
	}
	for i := 0; i < conn; i++ { // Wait for active workers.
		if url := <-finished; url != "" {
			results[url]++
		}
	}
	return
}

func worker(url *string, finished chan<- string, client *http.Client) {
	resp, err := client.Get(*url)
	if err != nil {
		finished <- *url
	} else {
		finished <- ""
	}
	if resp != nil {
		resp.Body.Close()
	}
}
