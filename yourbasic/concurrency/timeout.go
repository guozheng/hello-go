package concurrency

import (
	"fmt"
	"time"
)

func Timeout() (n int) {
	ch := make(chan int)
	go func() {
		select {
		case n := <-ch:
			fmt.Printf("got value from channel: %v\n", n)
		case <-time.After(2 * time.Second):
			fmt.Println("time out after 2 seconds")
			ch <- 10
		}
	}()
	time.Sleep(4 * time.Second) // wait for timeout of 2 second
	ret := <-ch
	fmt.Printf("done, got value: %d\n", ret)
	return ret
}

// time.After underlying time.Timer will not be GCed if it is not triggered
// so we can create a new timer instead and call Stop method
func TimeoutByTimer() (n int) {
	ch := make(chan int)
	timer := time.NewTimer(2 * time.Second)
	defer timer.Stop()
	go func() {
		select {
		case n := <-ch:
			fmt.Printf("got value from channel: %v\n", n)
		case <-timer.C:
			fmt.Println("timeout after 2 seconds")
			ch <- 10
		}
	}()
	time.Sleep(4 * time.Second) // wait for timeout of 2 seconds
	ret := <-ch
	fmt.Printf("done, got value: %d\n", ret)
	return ret
}

func RepeatedRun() {
	n := 0
	for now := range time.Tick(time.Second) {
		if n == 5 {
			fmt.Println("DONE")
			return
		}
		n++
		fmt.Printf("%d - %v\n", n, now)
	}
}

func ScheduledRun() {
	timer := time.AfterFunc(time.Second, func() {
		fmt.Println("after waiting for 1 second, DONE")
	})
	defer timer.Stop()
	time.Sleep(2 * time.Second) // wait for the scheduled run to trigger
}

