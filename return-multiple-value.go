package main

import "fmt"

func getFullName() (string, string, string) {
	return "Muhammad", "Noval", "Supriyadi"
}

func main() {
	firstname, middleName, _ := getFullName()
	fmt.Println(firstname, middleName)
}
