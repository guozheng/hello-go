package concurrency

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestSequentialRacer(t *testing.T) {
	slowServer := mockServer(20 * time.Millisecond)
	defer slowServer.Close()

	fastServer := mockServer(0)
	defer fastServer.Close()

	want := fastServer.URL
	got := SequentialRacer(slowServer.URL, fastServer.URL)

	if want != got {
		t.Errorf("got %q, want %q", got, want)
	}
}

func mockServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}
