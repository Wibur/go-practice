package main

import (
	"golang.org/x/tour/pic"
)

func Pic(dx, dy int) [][]uint8 {
	var pic [][]uint8
	for i := 0; i < dy; i++ {
		var one_line []uint8
		for j := 0; j < dx; j++ {
			// 後方公式 會生成不一樣圖形 (x+y)/2, x*y, x^y, x*log(y) 和 x%(y+1)
			one_line = append(one_line, uint8(i*j))
		}
		pic = append(pic, one_line)
	}
	return pic
}

func main() {
	pic.Show(Pic)
}