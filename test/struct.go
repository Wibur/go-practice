package main

import "fmt"

// 為了代碼中 多層結構資料封裝
type person struct {
	realName string
	age      int
	hobby    []string
	gender   string
}

func main() {
	var robert person
	robert.realName = "robert Kuo"
	robert.age = 999
	robert.hobby = []string{"game", "basketball", "read book"}

	fmt.Printf("%T, %v", robert, robert)
}
