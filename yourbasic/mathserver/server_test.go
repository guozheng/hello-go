package mathserver_test

import (
	"fmt"
	"github.com/guozheng/hello-go/yourbasic/mathserver"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestServerDivideByZero(t *testing.T) {
	s := mathserver.MathServer()

	divideByZero := &mathserver.Work{
		Op:    func(a, b int) int { return a / b },
		A:     100,
		B:     0,
		Reply: make(chan int),
	}
	s <- divideByZero

	select {
	case res := <-divideByZero.Reply:
		fmt.Println(res)
	case <-time.After(time.Second):
		fmt.Println("No result in one second.")
	}
	// Output: No result in one second.
}

func TestServerSum(t *testing.T) {
	s := mathserver.MathServer()

	sum := &mathserver.Work{
		Op:    func(a, b int) int { return a + b },
		A:     100,
		B:     50,
		Reply: make(chan int),
	}
	s <- sum

	select {
	case res := <-sum.Reply:
		fmt.Println(res)
	case <-time.After(time.Second):
		fmt.Println("No result in one second.")
	}
	// Output: 200
}