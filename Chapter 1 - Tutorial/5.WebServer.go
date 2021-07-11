// Simple echo webserver that counts the number of requests so far

package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

var mu sync.Mutex
var count int

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/count", counter)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	mu.Lock()
	count++
	mu.Unlock()
	fmt.Println("Handler", r)
	fmt.Fprintf(w, "URL: %q\n", r.URL.Path)
}

func counter(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Counter")
	mu.Lock()
	fmt.Fprintf(w, "Count: %d\n", count)
	mu.Unlock()
}
