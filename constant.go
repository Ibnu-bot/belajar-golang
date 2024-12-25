package main

import "fmt"

func main() {
	const (
		firstName string = "Ibnu"
		lastName         = "Hasan"
	)
	fmt.Println(firstName)
	fmt.Println(lastName)

	//Error
	// firstName ="woy"
	// lastName ="hai"
}