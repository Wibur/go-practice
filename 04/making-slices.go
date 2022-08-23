package main

import "fmt"

func main() {
	// make 用來初始化 特別的型別
	// make([] , len, cap)
	a := make([]int, 5)
	printSlices("a", a)

	b := make([]int, 0, 5)
	printSlices("b", b)

	c := b[:2]
	printSlices("c", c)

	d := c[2:5] // b的容量給5 因此最多就到5 超過 就會報錯
	printSlices("d", d)

	initData := []int{1, 3, 5, 7, 9}
	e := make([]int, 5, 15)
	copy(e, initData) // 把slice複製到另一個slice
	printSlices("e", e)
}

func printSlices(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n", s, len(x), cap(x), x)
}
