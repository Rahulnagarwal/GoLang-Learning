package main

import "fmt"

func modifyFn(val *int) {
	*val += 10 // Modifying the value at the address pointed to by val
}

func main() {
	num := 10
	ptr := &num // ptr holds the address of num

	fmt.Println("Value of num:", num)
	fmt.Println("Address of num:", ptr)
	fmt.Println("Value at the address stored in ptr:", *ptr) // Dereferencing ptr to get the value of num

	var ptr2 *int // Declaring a pointer to an int
	if ptr2 == nil {
		fmt.Println("ptr2 is nil, it does not point to any address.")
	}

	// **** accessing and modifying the value through the pointer ***
	value := 20
	modifyFn(&value) // Passing the address of value to the function
	fmt.Println("Modified value:", value)
}
