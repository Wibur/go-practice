package main

import (
	"fmt"
	"sync"
	"time"
)

type SafeNumber struct {
	v   int
	mux sync.Mutex
}

// routine 資源共享 race condition
func main() {
	total := SafeNumber{v: 0}
	for i := 0; i < 1000; i++ {
		go func() {
			total.mux.Lock()
			total.v++
			total.mux.Unlock()
		}()
	}
	time.Sleep(time.Second)
	total.mux.Lock()
	fmt.Println(total.v)
	total.mux.Unlock()
}

func say(s string) {
	start := time.Now()
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
	elapsed := time.Since(start)
	fmt.Println(elapsed)

}
