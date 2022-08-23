package main

import "fmt"

func main() {
	// strangeRange()
	praAssignment()
}

func strangeRange() {
	v := []int{1, 2, 3}
	for i := range v {
		v = append(v, i)
	}
	fmt.Println(v)
}

func praAssignment() {
	a := [3]int{1, 2, 3}
	b := a
	fmt.Println("[before]&a: %p,a: %v; &b: %v\n", &a[0], a, &b[0], b)
	b[0] = 6
	fmt.Println("[after]&a: %p,a: %v; &b: %v\n", &a, a, &b, b)
}
