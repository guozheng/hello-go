package iteration

func Repeat(c string, i int) string {
	var s string
	for j := 0; j < i; j++ {
		s += c
	}
	return s
}
