package main

import "fmt"

func endApp() {
	message := recover()
	if message != nil {
		fmt.Println("Error dengan message: ", message)
	}
	fmt.Println("APLIKASI SELESAI")
}

func runApp(error bool) {
	defer endApp()
	if error {
		panic(" APLIKASI ERROR")
	}
	fmt.Println("APLIKASI JALAN")
}

func main() {
	runApp(false)
}
