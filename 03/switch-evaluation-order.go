package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("When's Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("today")
	case today + 1:
		fmt.Println("Tomorrow")
	case today + 2:
		fmt.Println("In Two days")
	default:
		fmt.Println("Too far day")
	}
}