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
	http.HandleFunc("/count", handler)
	fmt.Println("Starting count server...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/count" {
		http.Error(w, fmt.Sprintf("path not supported: %v", r.URL.Path), http.StatusNotFound)
		return
	}

	switch r.Method {
	case "GET":
		mu.Lock()
		fmt.Fprintf(w, "Count: %d\n", count)
		log.Printf("Show current count: %d\n", count)
		mu.Unlock()
	case "POST":
		mu.Lock()
		count++
		fmt.Fprintf(w, "Count updated")
		log.Printf("Current count updated")
		mu.Unlock()
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

