package main

import "fmt"

type Filter func(string) string //type declaration

func sayHelloWithFilter(name string, filter Filter) {
	nameFiltered := filter(name)
	fmt.Println("Hello ", nameFiltered)
}

func spamFilter(name string) string {
	if name == "Babi" {
		return "..."
	} else {
		return name
	}
}

func main() {
	sayHelloWithFilter("Noval", spamFilter)
	sayHelloWithFilter("Babi", spamFilter)
}
