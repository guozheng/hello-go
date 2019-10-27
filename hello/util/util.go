package util

import (
	"errors"
	"fmt"
	"net"
	"runtime"
	"time"
)

func init() {
	fmt.Println("util init...")
}

// GetLocalIPAddress gives local ip address of running machine
func GetLocalIPAddress() (string, error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return "", err
	}
	for _, a := range addrs {
		if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String(), nil
			}
		}
	}
	return "", errors.New("no ipv4 address found")

}

// GetOs gives local machine's OS
func GetOs() string {
	switch os := runtime.GOOS; os {
	case "darwin":
		return "Mac OSX"
	case "linux":
		return "Linux"
	default:
		return "unknown"
	}
}

// RunTasks executes tasks in parallel
func RunTasks(c1, c2 chan<- string) {
	go speakAndSleep(1, 4, c1)
	go speakAndSleep(2, 2, c2)
}

func speakAndSleep(num, sec int, c chan<- string) {
	fmt.Printf("task %d, sleep %d sec\n", num, sec)
	time.Sleep(time.Duration(sec) * time.Second)
	fmt.Printf("task %d, wake up after %d sec\n", num, sec)
	c <- fmt.Sprintf("task %d DONE ", num)
}
