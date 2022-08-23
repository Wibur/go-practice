package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	findString := "how do you do"
	// 字串切割
	s := strings.Split(findString, " ") // string[]
	result := make(map[string]int, 10)
	for _, v := range s {
		if _, ok := result[v]; !ok {
			result[v] = 1
		} else {
			result[v]++
		}
	}
	// result["how"] = strings.Count(findString, "how")
	// result["do"] = strings.Count(findString, "do")
	// result["you"] = strings.Count(findString, "you")
	fmt.Println(result)
	// 找中文字
	var count int = 0
	chineseString := "123你好嗎"
	for _, v := range chineseString {
		if unicode.Is(unicode.Han, v) {
			count++
		}
	}
	fmt.Println(count)
}
