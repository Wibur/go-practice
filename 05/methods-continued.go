package main

import (
	"fmt"
	"math"
)

type MyFloat float64

// 也可以使用非結構體類型的聲明方法
// 注意： 接收者的類型定義和方法宣告 必須在同一包內
func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func main() {
	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())
}