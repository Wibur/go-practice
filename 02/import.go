package main

import (
	"fmt"
	"math"
)

func main() {
	// %g	%e for large exponents, %f otherwise.
	// Nextafter 回傳x之後到y的下一個 可顯示的float64值
	fmt.Println("Now u have %g problems.", math.Nextafter(2, 3))
}