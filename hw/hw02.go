package main

import (
	"fmt"
)

// 自定義類型
type person struct {
	age int
	weight int
	hobby []string
	gender string
	name string
}

/**
名字中包含1個 'e' or 'E' 分一個金幣
名字中包含1個 'i' or 'I' 分兩個金幣
名字中包含1個 'o' or 'O' 分三個金幣
名字中包含1個 'u' or 'U' 分四個金幣
*/
var (
	coins  = 5000
	names  = [5]string{"vancer", "robert", "damein", "leon", "allen"}
	result = make(map[string]int, len(names))
)

func main() {
	// fmt.Println(result)
	for _, name := range names {
		result[name] = 0
		disPatchCoin(name)
	}
	fmt.Println("all:", result, "left:", coins)
}

func disPatchCoin(name string) {
	for _, v := range name {
		switch v {
		case 'e', 'E':
			result[name]++
			coins--
		case 'i', 'I':
			result[name] += 2
			coins -= 2
		case 'o', 'O':
			result[name] += 3
			coins -= 3
		case 'u', 'U':
			result[name] += 4
			coins -= 3
		}
	}
}
