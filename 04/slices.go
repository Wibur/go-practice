package main

import "fmt"

func main() {
	primes := [6]int{2,3,5,7,11,13}
	// a[low : high] 下面為取 index 1~4的意思
	var s []int = primes[1:4]
	fmt.Println(s)
}