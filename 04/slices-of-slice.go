package main

import (
	"fmt"
	"strings"
)

func main() {
	// 創建井字板
	board := [][]string{
		[]string{"_","_","_"},
		[]string{"_","_","_"},
		[]string{"_","_","_"},
	}
	// 圈圈叉叉
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for i := 0; i < len(board); i++ {
		// strings.Join 把字符串用指定符號相接
		fmt.Printf("%s\n", strings.Join(board[i], " "))
	}
}