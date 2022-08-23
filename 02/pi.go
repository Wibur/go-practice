package main

import (
	"fmt"
	"math"
)

func main() {
	// pi沒被導出 因此math.pi會報錯
	fmt.Println(math.Pi)
}
