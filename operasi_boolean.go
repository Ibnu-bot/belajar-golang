package main

import "fmt"

func main() {

	//Operasi boolean && (dan)
	var nilaiAkhir = 90
	var absensi = 80

	var lulusNilaiAkhir bool = nilaiAkhir > 80
	var lulusAbsensi bool = absensi > 80

	var lulus bool = lulusNilaiAkhir && lulusAbsensi

	fmt.Println(lulus)

	//Operasi boolean || (atau)
	var lulusor bool = lulusNilaiAkhir || lulusAbsensi
	fmt.Println(lulusor)

	// Operasi boolean ! (not)
	var tidakLulus bool = !lulus
	fmt.Println(tidakLulus)
}