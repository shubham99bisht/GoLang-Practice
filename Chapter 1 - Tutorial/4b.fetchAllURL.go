/*
Fetch All URL fetches all the URLs passed via cmd line arguments concurrently using goroutines
To run: $ go run 4b.fetchAllURL.go http://golang.org http://gopl.io http://godocs.org
*/

package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func main() {
	start := time.Now()
	ch := make(chan string) // Make a string channel
	for _, url := range os.Args[1:] {
		go fetch(url, ch) // Start a goroutine
	}

	for range os.Args[1:] {
		fmt.Println(<-ch) // Receive from channel ch
	}
	fmt.Println(time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}
	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)
}
