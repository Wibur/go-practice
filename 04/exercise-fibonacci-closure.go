package main

import "fmt"
// 0 1 1 2 3
func fibonacci() func() int {
	f1, f2 := 0, 1
	return func() int {
		f := f1 			//  0     1   1    2    3
		f1, f2 = f2, f2 + f // 1,1  1,2  2,3  3,5  5,8
		return f
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}