package main

import "fmt"

func main() {
	pow := make([]int, 10)
	// 只需要索引 直接忽略第二個變數就可
	for i := range pow {
		pow[i] = 1 << uint(i) // == 2 ** i
	}
	for _, v := range pow {
		fmt.Printf("%d\n", v)
	}
}