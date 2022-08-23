package main

import "fmt"

type USB interface {
	Name() string
	Connector
}

type Connector interface {
	Connect()
}

type PhoneConnecter struct {
	name string
}

type TVConnecter struct {
	name string
}

func (TV TVConnecter) Connect() {
	fmt.Println("connect: ", TV.name)
}

func (TV TVConnecter) Name() string {
	return TV.name
}

func (pc PhoneConnecter) Name() string {
	return pc.name
}

func (pc PhoneConnecter) Connect() {
	fmt.Println("Connection:", pc.name)
}

func main() {
	a := TVConnecter{"vancer"}
	a.Connect()
	a.name = "TV Connector"
	a.Connect()
}

func DisConnect(usb USB) {
	switch v := usb.(type) {
	case PhoneConnecter:
		fmt.Println("Disconnect", v.name)
	default:
		fmt.Println("UnKnown")
	}

}
