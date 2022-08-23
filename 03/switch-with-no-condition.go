package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()
	// 不給條件 相當於true
	switch {
	case t.Hour() < 12:
		fmt.Println("Good morning!")
	case t.Hour() < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}