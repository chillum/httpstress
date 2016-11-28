/*
Package httpstress is a library for HTTP stress testing

It launches one goroutine per concurrent connection and does not count successful attempts.

It follows HTTP redirects. Non-200 HTTP return codes are considered as errors.

A CLI utility is avaliable at github.com/chillum/httpstress
*/
package httpstress

import (
	"errors"
	"net/http"
	"strings"
)

// Version is the library version
const Version = "2.1.1"

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
	if conn < 1 {
		err = errors.New("connections number cannot be less than 1")
		return
	}

	for _, i := range urls {
		if !strings.HasPrefix(i, "http://") && !strings.HasPrefix(i, "https://") {
			err = errors.New("not a HTTP(S) URL: " + i)
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
	client.CheckRedirect = redirect
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

// Get the URL and report status to the channel.
func worker(url *string, finished chan<- string, client *http.Client) {
	var err error
	var resp *http.Response
	var req *http.Request

	req, err = http.NewRequest("GET", *url, nil)
	if err != nil {
		finished <- *url
	} else {
		req.Header.Set("User-Agent", "httpstress")
		resp, err = client.Do(req)
		if err == nil {
			if resp.StatusCode == 200 { // Check status code.
				finished <- ""
			} else {
				finished <- *url
			}
		} else {
			finished <- *url
		}
	}
	if resp != nil {
		resp.Body.Close()
	}
}

// Check HTTP redirects.
func redirect(req *http.Request, via []*http.Request) error {
	// When redirects number > 10 probably there's a problem.
	if len(via) >= 10 {
		return errors.New("stopped after 10 redirects")
	}
	// Redirects don't get User-Agent.
	req.Header.Set("User-Agent", "httpstress")
	return nil
}
