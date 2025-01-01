package main

import "fmt"

func main() {
	var names [3]string
	names[0] = "Hallo"
	names[1] = "Saya"
	names[2] = "Ibnu"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	var values = [...] int {
		90,
		80,
		84,
		100,
		120,
	}
	fmt.Println(values)
	fmt.Println(len(values))

}