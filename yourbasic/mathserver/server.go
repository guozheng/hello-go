package mathserver

import "log"

// MathServer creates a new mathserver that accepts Work requests
// through the req channel.
func MathServer() (req chan<- *Work) {
	wc := make(chan *Work)
	go serve(wc)
	return wc
}

type Work struct {
	Op    func(int, int) int
	A, B  int
	Reply chan int // Server sends result on this channel.
}

func serve(wc <-chan *Work) {
	for w := range wc {
		go safelyDo(w)
	}
}

func safelyDo(w *Work) {
	// Regain control of a panicking goroutine to avoid
	// killing the other executing goroutines.
	defer func() {
		if err := recover(); err != nil {
			log.Println("work failed:", err)
		}
	}()
	do(w)
}

func do(w *Work) {
	w.Reply <- w.Op(w.A, w.B)
}