package concurrency

import (
	"fmt"
	"testing"
)

func TestRace(t *testing.T)  {
	n := Race()
	if *n == 2 {
		t.Errorf("got %v, should NOT be 2", *n)
	}
	fmt.Println("got 2")
}

func TestNoRace(t *testing.T) {
	n := NoRace()
	if n != 2 {
		t.Errorf("got %v, should be 2", n)
	}
	fmt.Println("got 2")
}
