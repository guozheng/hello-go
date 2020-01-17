package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"
)

func main()  {
	start := time.Now()
	urls := [2]string{
		"http://www.google.com",
		"http://www.yahoo.com",
	}
	ch := make(chan string)
	for _, url := range urls {
		go fetch(url, ch)
	}
	select {
	case content := <-ch:
		fmt.Printf("Got a response: %s\n", content)
		return
	case <-time.After(10 * time.Second):
		fmt.Errorf("Timed out")
		return
	}

	fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
}

func fetch(url string, ch chan<- string)  {
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}

	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("time: %.2fs, %7d bytes, url: %s", secs, nbytes, url)
}
