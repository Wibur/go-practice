package main

import "fmt"

func main() {
	names := []int{1,2,3,4,5,6,7,8,9,10,11,12}
	// 切片的下界 默認為0
	a := names[:10]
	fmt.Println(a)
	// 上界為切片長度
	b := names[:]
	fmt.Println(b)
}