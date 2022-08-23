package main

import "fmt"

func main() {
	var i, j int = 1, 2
	// := 結構不能使用在函式外
	k := 3
	c, python, java := true, false, "yo"

	fmt.Println(i, j, k, c, python, java)
}