package main

import "fmt"

func main() {
	fmt.Println("Learning array in golang")

	var name [5]string
	name[0] = "John"
	name[2] = "Doe"

	fmt.Println("names are :", name)

	var numbers = [6]int{1, 2, 3, 4, 5}
	fmt.Println("numbers are :", numbers)

	fmt.Println("length of number array is :", len(numbers))
}
