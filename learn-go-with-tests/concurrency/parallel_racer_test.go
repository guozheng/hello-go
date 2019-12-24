package concurrency

import (
	"testing"
	"time"
)

func TestParallelRacer(t *testing.T) {
	slowServer := mockServer(20 * time.Millisecond)
	defer slowServer.Close()

	fastServer := mockServer(19 * time.Millisecond)
	defer fastServer.Close()

	want := fastServer.URL
	got, err := ParallelRacer(slowServer.URL, fastServer.URL)

	if err != nil {
		t.Error("not expecting an error but got one")
	}

	if want != got {
		t.Errorf("got %q, want %q", got, want)
	}
}