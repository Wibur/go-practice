package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}
// go 沒有類別 但可以為結構體定義方法
// Abs 方法拥有一个名为 v，类型为 Vertex 的接收者。
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
}