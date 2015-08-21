package main

import (
	"errors"
	"net/http"
	"regexp"
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
		if m, _ := regexp.MatchString("^https?://", i); !m {
			err = errors.New("Not a HTTP/HTTPS URL: " + i)
			return
		}
	}

	results = make(map[string]int)
	failures := make(chan string)
	finished := make(chan bool)
	total := len(urls) - 1
	trans := &http.Transport{MaxIdleConnsPerHost: conn} // Use persistent connections.
	client := &http.Client{Transport: trans}
	n := 0
	i := 0

	// Ensure, every URL gets a request.
	if max < len(urls) {
		max = len(urls)
	}

	go logger(failures, results)
	for ; i < conn; i++ { // Launch initial workers.
		go worker(&urls[n], failures, finished, client)

		if n < total {
			n++
		} else {
			n = 0
		}
	}
	for ; i < max; i++ { // Launch more workers as initial finish.
		if <-finished {
			go worker(&urls[n], failures, finished, client)

			if n < total {
				n++
			} else {
				n = 0
			}
		}
	}
	for i := 0; i < conn; i++ { // Wait for active workers.
		<-finished
	}
	return
}

func logger(failures <-chan string, results map[string]int) {
	for {
		select {
		case url := <-failures:
			results[url]++
		}
	}
}

func worker(url *string, failures chan<- string, finished chan<- bool, client *http.Client) {
	resp, err := client.Get(*url)
	if err != nil {
		failures <- *url
	}
	if resp != nil {
		resp.Body.Close()
	}
	finished <- true
}
