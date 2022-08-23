package main

import "fmt"

func adder() func(int) int{
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	// 返回一個閉包
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			pos(i),
			neg(-2*i),
		)
	}
}