package concurrency

func Race() (m *int) {
	wait := make(chan struct{})
	var p = new(int)
	*p = 0
	go func(q *int) {
		*q++;
		close(wait)
	}(p)
	*p++ // race condition
	<-wait // block
	return p
}

func NoRace() (m int)  {
	ch := make(chan int)
	go func() {
		n := 0
		n++
		ch <- n
	}()
	n := <-ch
	n++
	return n
}
