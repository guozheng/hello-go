package concurrency

import (
	"fmt"
	"sync"
	"time"
)

func RunInGroup()  {
	var group sync.WaitGroup
	group.Add(2)
	go func() {
		time.Sleep(time.Second)
		fmt.Println("Task 1 done")
		group.Done()
	}()
	go func() {
		time.Sleep(2 * time.Second)
		fmt.Println("Task 2 done")
		group.Done()
	}()
	group.Wait() // wait till all tasks are done
}
