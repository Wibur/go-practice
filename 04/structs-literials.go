package main

import "fmt"

type Vertex struct {
	X, Y int
}

var (
	v1 = Vertex{1, 2}	//	創建一個vertex 類型的結構體
	v2 = Vertex{X: 1}	//	Y: 0 被默認賦予
	v3 = Vertex{}		//  X: 0 Y: 0
	p = &Vertex{1, 2}	//  創建一個結構指針
)

func main() {
	fmt.Println(v1, p, v2, v3)
}