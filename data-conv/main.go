package main

import (
	"fmt"
	"strconv"
)

// strconv , Atoi for Integer to string, Iton for interger to string.
func main() {
	num := 42
	fmt.Println("The number is:", num)
	fmt.Printf("Type of num is %T\n", num)

	var data string = strconv.Itoa(num) // Type conversion from int to float64
	fmt.Println("Data as string:", data)
	fmt.Printf("Type of data is %T\n", data)

	number_string := "123"
	converted_number, _ := strconv.Atoi(number_string) // Type conversion from string to int

	fmt.Println("Converted number:", converted_number)
	fmt.Printf("Type of converted_number is %T\n", converted_number)
}
