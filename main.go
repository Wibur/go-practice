package main

import (
	"fmt"
	blc "vancer/blc"
)

func main() {
	block := blc.CreateBlock("Vancer Block", 1, []byte{0,0,0,0,0,0,0,0,0,0})
	fmt.Println(block)
}