package main

import "fmt"

func main() {
	i := 10
	ip := &i
	fmt.Printf("原指針內存地址: %p\n", &ip)
	modify(ip)
	fmt.Println("int值被修改了,新值為:", i)
}

func modify(ip *int) { // 帶進來的參數 是複製出來新的位址 但值不變
	fmt.Printf("函數裡收到的指針地址: %p\n", &ip)
	*ip = 1
}