package util_test

import (
	"fmt"
	"testing"

	"github.com/guozheng/hello-go/util"
)

func TestGetOs(t *testing.T) {
	expected := "Mac OSX"
	if actual := util.GetOs(); actual != expected {
		t.Errorf("expected: %q, actual: %q, time to get a Mac", expected, actual)
	}
}

func ExampleTestGetOs() {
	fmt.Printf("OS: %s", util.GetOs())
	// Output: OS: Mac OSX
}

func BenchmarkGetIPAddress(b *testing.B) {
	for i := 0; i < b.N; i++ {
		os, _ := util.GetLocalIPAddress()
		fmt.Println(os)
	}
}
