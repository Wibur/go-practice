package main

import "fmt"

func main() {
	m := make(map[string]int)
	// m[key] = elem

	m["Answer"] = 42
	fmt.Println("The val:", m["Answer"])
	
	m["Answer"] = 48
	fmt.Println("The val:", m["Answer"])
	// 刪除元素
	delete(m, "Answer")
	fmt.Println("The val:", m["Answer"])
	// elem, ok := m[key] 若 elem 或 ok 还未声明，你可以使用短变量声明
	v, ok := m["Answer"]
	fmt.Println("The val:", v, "Present?", ok)
}