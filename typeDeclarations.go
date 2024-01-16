package main 

import "fmt"

func main() {
	type NoKTP string
	type Married bool

	var noKtpNoval NoKTP = "3827173982328932"
	var marriedStatus Married = true
	fmt.Println(noKtpNoval)
	fmt.Println(marriedStatus)
}