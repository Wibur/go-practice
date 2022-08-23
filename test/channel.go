package main

import (
	"fmt"
	"math/rand"
	"time"
)

func longTimeRequest(r chan<- int32) {
	time.Sleep(time.Second * 3) // 模拟一个工作负载
	r <- rand.Int31n(100)
}

func sumSquares(a, b int32) int32 {
	return a*a + b*b
}

func main() {
	rand.Seed(time.Now().UnixNano())

	result := make(chan int32, 2)
	go longTimeRequest(result)
	go longTimeRequest(result)

	fmt.Println(sumSquares(<-result, <-result))
}
