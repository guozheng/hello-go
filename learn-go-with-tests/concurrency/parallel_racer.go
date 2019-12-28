package concurrency

import (
	"fmt"
	"net/http"
	"time"
)

func ParallelRacer(url1, url2 string) (winner string, err error)  {
	select {
	case <-fetch(url1):
		return url1, nil
	case <-fetch(url2):
		return url2, nil
	case <-time.After(25 * time.Second):
		return "", fmt.Errorf("timeout error after waiting for %s and %s", url1, url2)
	}
}

func fetch(url string) chan struct{} {
	ch := make(chan struct{})
	go func() {
		http.Get(url)
		close(ch)
	}()
	return ch
}
