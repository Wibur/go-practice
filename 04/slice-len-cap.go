package main

import "fmt"

func main() {
	s := []int{1,2,3,4,5}
	// 切片長度 => 所包含的元素個數
	// 切片容量 => 從第一個元素到最後一個元素 共有幾個元素
	s = s[:0]
	printSlices(s)	// 0, 5(0開始數)
	
	s = s[:4]
	printSlices(s)	// 4, 5(0開始數)

	s = s[2:]
	printSlices(s)	// 2, 3(2開始數)
}

func printSlices(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}