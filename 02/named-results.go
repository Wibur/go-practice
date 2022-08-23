package main

import "fmt"

func split(s int) (x, y int) {
	x = s * 4 % 9
	y = s - x
	return
}

func main() {
	fmt.Println(split(17))
}