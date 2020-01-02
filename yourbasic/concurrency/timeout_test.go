package concurrency

import "testing"

func TestTimeout(t *testing.T) {
	n := Timeout()
	if n != 10 {
		t.Errorf("did not get expected value 10, got %d", n)
	}
}

func TestTimeoutByTimer(t *testing.T) {
	n := TimeoutByTimer()
	if n != 10 {
		t.Errorf("did not get expected value 10, got %d", n)
	}
}

func TestRepeatedRun(t *testing.T) {
	RepeatedRun()
}

func TestScheduledRun(t *testing.T) {
	ScheduledRun()
}