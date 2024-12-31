package main

import "fmt"

func main() {

	//Perbandingan String
	var name1 = "Ibnu"
	var name2 = "Ibnu"

	var result bool = name1 == name2
	fmt.Println("Perbandingan string:", result)

	//Perbandingan Number
	var num1 = 10
	var num2 = 20

	fmt.Println("num1 == num2:", num1 == num2)
	fmt.Println("num1 != num2:", num1 != num2)
	fmt.Println("num1 < num2:", num1 < num2)
	fmt.Println("num1 > num2:", num1 > num2)
	fmt.Println("num1 <= num2:", num1 <= num2)
	fmt.Println("num1 >= num2:", num1 >= num2)

	
}