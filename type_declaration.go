package main

import "fmt"

func main() {

	type NoKTP string
	var ktpIbnu NoKTP = "123456"

	var contoh string = "22222"
	var contohKTP NoKTP = NoKTP(contoh)

	fmt.Println(ktpIbnu)
	fmt.Println(contohKTP)

}