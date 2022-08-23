package main

import "fmt"

type TZ int

func (a *TZ) Increase(num int) {
	*a += TZ(num)
}

func main() {
	var x TZ
	x.Increase(100)
	fmt.Println(x)
}
