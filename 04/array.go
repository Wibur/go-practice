package main

import "fmt"

func main() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "world"
	fmt.Println(a[0], a[1])
	fmt.Println(a)
	// 陣列的簡潔使用
	primes := [6]int{2,3,4,5,6,7}
	fmt.Println(primes)
}