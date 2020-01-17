package string

import "testing"

//BenchmarkStringGen          	1000000000	         0.00126 ns/op
//BenchmarkStringGenBuilder   	1000000000	         0.000092 ns/op
//BenchmarkStringGenBuffer    	1000000000	         0.000046 ns/op

func BenchmarkStringGen(b *testing.B) {
	StringGen(1000)
}

func BenchmarkStringGenBuilder(b *testing.B) {
	StringGenBuilder(1000)
}

func BenchmarkStringGenBuffer(b *testing.B) {
	StringGenBuffer(1000)
}