package main 

import "fmt"

func main() {
	var nilaiAkhir = 80
	var absensi = 75

	var lulusNilaiAkhir bool = nilaiAkhir >= 80
	var lulusAbsensi bool = absensi >= 80
	
	fmt.Println(lulusNilaiAkhir)
	fmt.Println(lulusAbsensi)

	var lulus bool = lulusNilaiAkhir && lulusAbsensi

	fmt.Println(lulus)

	fmt.Println(nilaiAkhir >= 80 && absensi >= 80)
}