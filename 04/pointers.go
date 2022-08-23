package main

import "fmt"

func main() {
	i, j := 42, 2701

	p := &i			// 指向i
	fmt.Println(*p)	// 通過指針讀取i的值
	*p = 21 		// 通過指針設置i的值
	fmt.Println(i)

	p = &j	
	*p = *p / 37
	fmt.Println(j)
}