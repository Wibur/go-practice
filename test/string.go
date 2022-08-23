package main

import (
	"fmt"
)

func main() {
	var out []*int
	for i := 0; i < 3; i++ {
		x := i // 如果使用i直接插入 i會因為傳遞位址都相同的原因 打印出來都是一樣的
		out = append(out, &x)
	}
	fmt.Println("Values:", *out[0], *out[1], *out[2])
	fmt.Println("Addresses:", out[0], out[1], out[2])
	test()
}

func test() {
	var out [][]int
	for _, i := range [][1]int{{1}, {2}, {3}} {
		out = append(out, i[:])
	}
	fmt.Println("Values:", out)
}
