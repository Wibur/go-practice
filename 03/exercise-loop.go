package main

import (
	"fmt"
	"math"
)

func Sqrt(x float64) float64 {
	z := float64(1)
	val := float64(0)
	for {
		z = z - (z * z - x ) / (2 * z)
		// 1e-10 => 1*10^-10 此處為了無限接近0
		if math.Abs(val - z) < 1e-10 {
			break
		}
		val = z
	}
	return val
}

func main() {
	fmt.Println(Sqrt(2))
}