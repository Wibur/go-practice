package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// 注意： 這個程式的運行環境是固定的，因此 rand.Intn 總是會返回相同的數字
	// rand.Seed(n) => 生成不同的種子數
	fmt.Println("My favorite number is", rand.Intn(10))
}