package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}
	p := &v	  // 將v的值 記載p的記憶體位址 因此其中一方改變 另一方也會變
	p.X = 1e9 // (*p).X的意思
	// v = Vertex{9, 2}
	fmt.Println(v)
}