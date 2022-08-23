package main

import "fmt"

func main() {
	var k map[string]int         // map 壹定要初始化才能用
	k = make(map[string]int, 10) // 容量盡量一開始就設置好
	k["vancer"] = 10
	k["david"] = 99
	fmt.Println(k)

	// v, ok := k["vancer"]
	// if ok {
	// 	fmt.Println(v)
	// } else {
	// 	fmt.Println("no exists")
	// }

	for k, v := range k {
		fmt.Println(k, v)
	}
	// 元素類型為map的切片
	k2 := make([]map[string]int, 10, 10) // 長度為0
	// 要對內部的map初始化 才能用
	k2[0] = make(map[string]int, 10)
	k2[0]["hot"] = 999
	fmt.Println(k2)

	// 值為切片類型的map
	var k3 = make(map[string][]int, 10)
	k3["dvas"] = []int{1, 2, 3}
	fmt.Println(k3)
}
