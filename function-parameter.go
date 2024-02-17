package main

import "fmt"

func sayHelloTo(firstName string, midName string, lastName string) {
	fmt.Println("Hello", firstName, midName, lastName)
}

func main() {
	firstName := "Noval"
	sayHelloTo(firstName, "Muhammad", "Supriyadi")
	sayHelloTo("Muhammad", "Farhan", "Syahputera")
}
