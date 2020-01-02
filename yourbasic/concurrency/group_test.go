package concurrency

import (
	"testing"
	"time"
)

func TestRunInGroup(t *testing.T) {
	start := time.Now()
	RunInGroup()
	elapsed := time.Now().Sub(start)
	if elapsed < 2 * time.Second {
		t.Errorf("tasks should take more than 2 seconds, actual elapsed time %v seconds", elapsed)
	}
}
