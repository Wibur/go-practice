package main

import "fmt"

// func add(x int, y int) int {
// 可縮寫為
func add(x, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(42, 13))
}