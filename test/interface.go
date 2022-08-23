package main

import (
	"fmt"
	"sort"
)

func main() {
	family := []struct {
		Name string
		Age  int
	}{{"Irene", 12}, {"Justin", 45}, {"Monica", 42}}

	sort.SliceStable(family, func(i, j int) bool {
		return family[i].Age < family[j].Age
	})

	fmt.Println(family)

	idx := sort.Search(len(family), func(i int) bool {
		return family[i].Age == 45
	})
	fmt.Println(idx)
	GuessingGame()
}

func GuessingGame() {
	var s string
	fmt.Printf("Pick an integer from 0 to 100.\n")
	answer := sort.Search(100, func(i int) bool {
		fmt.Printf("Is your number <= %d? ", i)
		fmt.Scanf("%s", &s)
		return s != "" && s[0] == 'y'
	})
	fmt.Printf("Your number is %d.\n", answer)
}
