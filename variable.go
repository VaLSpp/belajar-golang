package main

import "fmt"

func main() {
	var name string

	name = "Muhammad Noval Supriyadi"
	fmt.Println(name)

	name = "Muhammad Noval"
	fmt.Println(name)

	// contoh yang menggunakan tipe data variable
	var age = 16
	fmt.Println(age)

	age = age + age + age
	fmt.Println(age)

	// alternatif jika tidak pakai kata kunci var
	umur := 16 // titik dua sama dengan (:=) cuman untuk deklarasi awalan saja
	fmt.Println(umur)

	umur = umur + umur + 8
	fmt.Println(umur)

	// yang menggunakan deklarasi multiple variable
	var (
		firstName = "Muhammad Noval"
		lastName  = "Supriyadi"
	)
	fmt.Println(firstName)
	fmt.Println(lastName)
}
