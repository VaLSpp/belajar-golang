package main

import "fmt"

func main() {

	var values = [...]int{
		90,
		80,
		95,
	}
	fmt.Println(values)
	fmt.Println(len(values))
	values [0] = 100
	fmt.Println(values)
}
