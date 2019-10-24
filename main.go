package main

import (
	"fmt"
	"time"

	"github.com/guozheng/hello-go/data"
	"github.com/guozheng/hello-go/util"
)

func main() {
	fmt.Println("hello, Go!")

	// use util methods
	localIP, err := util.GetLocalIPAddress()
	if err != nil {
		panic(err)
	}
	fmt.Printf("your ip: %s\n", localIP)
	localOs := util.GetOs()
	fmt.Printf("your os: %s\n", localOs)

	// use struct
	p := data.Person{Name: "Alex", Age: 10, Gender: "M"}
	fmt.Printf("p: %#v\n", p)
	ptr := &p
	ptr.Describe()
	ptr.SetName("Claire")
	ptr.SetAge(6)
	ptr.SetGender("F")
	ptr.Describe()

	// use go routine and channel
	c1 := make(chan string)
	c2 := make(chan string)
	util.RunTasks(c1, c2)
	sec := 5
	fmt.Printf("main sleep %d sec\n", sec)
	time.Sleep(time.Duration(sec) * time.Second)
	fmt.Printf("main wake up after %d sec\n", sec)
	fmt.Printf("task1 output: %s\n", <-c1)
	fmt.Printf("task2 output: %s\n", <-c2)
}
