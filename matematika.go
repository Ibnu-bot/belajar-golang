package main

import "fmt"

func main() {

	//Operasi matematika (+ - / * %modulus)
	var a = 100
	var b = 100
	var c = 5
	var d = 10
	var e = 50
	var f = 9
	var g = a + b*c / d - e % f
	fmt.Println(g)

	//Augmented Assignments
	var i = 10
	i += 10 // i = i + 10
	fmt.Println(i)
	
	i += 5 // i = i + 5
	fmt.Println(i)

	//Unary Operator
	var j = 1
	j++ // j = j + 1
	fmt.Println(j)

	j++ // j = j + 1
	fmt.Println(j)

	j-- // j = j - 1
	fmt.Println(j)

	j-- // j = j - 1
	fmt.Println(j)


}