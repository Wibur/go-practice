package main

import "fmt"

func main() {
	// [6]int{2,3,5,7,11,13} 下面那樣寫也可以
	q := []int{2,3,5,7,11,13}
	fmt.Println(q)

	r := []bool{true, false, true, false, true, false}
	fmt.Println(r)

	s := []struct{
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(s)
}