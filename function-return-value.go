package main

import "fmt"

func getHello(name string) string {
	if name == "" {
		return "Hello bang"
	} else {
		return "Hello " + name
	}
}

func main() {
	result := getHello("Noval")
	fmt.Println(result)

	fmt.Println(getHello(""))
}
