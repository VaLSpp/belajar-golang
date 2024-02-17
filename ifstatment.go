package main

import "fmt"

func main() {
	var name = "Supriyadi"

	if name == "Noval" {
		fmt.Println("Hello Noval")
	} else if name == "Noval" {
		fmt.Println("Hi, Noval")
	} else if name == "Supriyadi" {
		fmt.Println("Hi, Supriyadi")
	} else {
		fmt.Println("Hi, Boleh kenalan?")
	}

	if length := len(name); length > 5 {
		fmt.Println("Terlalu Panjang")
	} else {
		fmt.Println("Nama Sudah Benar")
	}
}
