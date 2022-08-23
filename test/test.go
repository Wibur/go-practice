package main

func main() {
	// var c chan<- bool
	var c1 chan bool
	// c<-false
	<-c1
}
