package string

import (
	"fmt"
	"strconv"
	"strings"
)

func StringGen(n int) (t string) {
	s := ""
	for i := 0; i < n; i++ {
		s = s + " " + strconv.Itoa(i)
	}
	return s
}

func StringGenBuilder(n int) (t string) {
	var b strings.Builder
	b.Grow(n * 2)
	for i := 0; i < n; i++ {
		fmt.Fprintf(&b, " ")
	}
	return b.String()
}

func StringGenBuffer(n int64) (t string) {
	buf := make([]byte, 0, n * 2)
	var i int64
	for i = 0; i < n; i++ {
		buf = strconv.AppendInt(buf, i, 10)
	}
	return string(buf)
}