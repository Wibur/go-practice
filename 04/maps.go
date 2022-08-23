package main

import (
	"fmt"
)

type Vertex struct {
	Lat, Long float64
}

type Vanessa struct {
	qq, yy string
}

var m map[string]Vertex
var t map[string]Vanessa

// 映射 將鍵映射到值
func main() {
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.68333, -74.39967,
	}
	fmt.Println(m["Bell Labs"])

	t = make(map[string]Vanessa)
	t["PHP"] = Vanessa{"ThinkPHP", "Laravel"}

	v, ok := t["PHP"]
	if !(ok == false) {
		fmt.Println(v)
	}
}