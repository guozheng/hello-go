package writer

import (
	"bytes"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGreet(t *testing.T) {
	buf := bytes.Buffer{}
	Greet(&buf, "Alex")
	got := buf.String()
	want := "Hi, Alex"
	assert.Equal(t, want, got, "want should match got")
}
