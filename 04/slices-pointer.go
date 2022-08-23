package main

import "fmt"

func main() {
	names := [4]string{
		"John",
		"Carter",
		"Vanessa",
		"George",
	}

	fmt.Println(names)

	a := names[0:2] // 取index 0-2
	b := names[1:3] // 取index 1-3
	fmt.Println(a, b)

	b[0] = "XXX"	// 更改切片的元素 底層陣列中對應的元素也會被一併更改
	fmt.Println(a, b)
	fmt.Println(names)
}
// 切片不儲存任何數據 只描述底層陣列中的一段