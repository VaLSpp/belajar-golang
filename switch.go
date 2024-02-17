package main

import "fmt"

//mirip dengan if else tapi versi sederhana nya

func main() {
	name := "Noval"

	switch name {
	case "Noval":
		fmt.Println("Hello Noval")
		fmt.Println("Hello Noval")
	case "Milki":
		fmt.Println("Hello Milki")
		fmt.Println("Hello Milki")
	default: //default itu sama aja dengan else
		fmt.Println("Hi, Boleh Kenalan?")
		fmt.Println("Hi, Boleh Kenalan?")
	}

	// switch length := len(name); length > 8 {
	// case true:
	// 	fmt.Println("Nama Terlalu Panjang")
	// case false:
	// 	fmt.Println("Nama Sudah Benar")
	// }

	length := len(name)
	switch {
	case length > 10:
		fmt.Println("Nama Terlalu Panjang")
	case length > 5:
		fmt.Println("Nama Lumayan Panjang")
	default:
		fmt.Println("Nama Sudah Benar")
	}
}
