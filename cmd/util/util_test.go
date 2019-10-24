package util_test

import (
	"fmt"
	"testing"

	"github.com/guozheng/hello-go/cmd/util"
)

func TestGetLocalIpAddress(t *testing.T) {
	expected := "192.168.1.97"
	if actual, _ := util.GetLocalIPAddress(); actual != expected {
		t.Errorf("expected: %q, actual: %q", expected, actual)
	}
}

func TestGetOs(t *testing.T) {
	expected := "Mac OSX"
	if actual := util.GetOs(); actual != expected {
		t.Errorf("expected: %q, actual: %q", expected, actual)
	}
}

func ExampleTestGetOs() {
	fmt.Printf("OS: %s", util.GetOs())
	// Output: OS: Mac OSX
}
