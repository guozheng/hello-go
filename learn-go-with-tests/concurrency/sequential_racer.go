package concurrency

import (
	"net/http"
	"time"
)

func SequentialRacer(url1, url2 string) (winner string)  {
	m1 := measureResponseTime(url1)
	m2 := measureResponseTime(url2)
	if m1 < m2 {
		return url1
	}

	return url2

}

func measureResponseTime(url string) time.Duration {
	start := time.Now()
	http.Get(url)
	return time.Since(start)
}
