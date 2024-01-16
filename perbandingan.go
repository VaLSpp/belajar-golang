package main

import "fmt"

func main() {
	var name1 = "noval"
	var name2 = "noval supriyadi"

	var result bool = name1 == name2 // == kalau sama true kalau berbeda false
	fmt.Println(result)

	var value1 = 100
	var value2 = 200

	fmt.Println(value1 > value2)
	fmt.Println(value1 < value2)
	fmt.Println(value1 == value2)
	fmt.Println(value1 != value2)
}
