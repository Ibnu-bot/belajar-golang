package main

import "fmt"

func main() {
	//Konversi Tipe data integer
	var nilai32 int32 = 32768
	var nilai64 int64 = int64(nilai32)
	var nilai16 int16 = int16(nilai32)

	fmt.Println(nilai32)
	fmt.Println(nilai64)
	fmt.Println(nilai16)

	//Konversi data karakter
	var name = "Ibnu Hasan"
	var i uint8 = name[0]
	var iString = string(i)

	fmt.Println(name)
	fmt.Println(i)
	fmt.Println(iString)
}