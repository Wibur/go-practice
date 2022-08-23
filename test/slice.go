package main

import "fmt"

func main() {
	s := []int{1, 2, 3, 4, 5}
	d := make([]int, 5, 5) // 長度5 容量5
	copy(d, s)             // copy 只會針對已初始化的元素改為源slice中對應的元素

	fmt.Printf("&d: %p, len: %d, cap: %d, d: %v\n", d, len(d), cap(d), d)
	// 切片就是一個框 框住了一塊連續的記憶體
	s2 := []int{} //len(s2)=0, cap(s2)=0, s2 != nil
	if s2 == nil {
		fmt.Println(1)
	}
	s3 := make([]int, 0) // len(s3)=0, cap(s3)=0, s3 1= nil
	if s3 != nil {
		fmt.Println(2)
	}
}
