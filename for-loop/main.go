package main

import "fmt"

func main() {
	// No bracket and can not start { from the next line.
	for i := 0; i < 5; i++ {
		fmt.Println("Numbers is", i)
	}

	// this is the range keyword that get index and value for array or slice and string.
	numbers := []int{39, 22, 43, 46, 85}
	for index, value := range numbers {
		fmt.Printf("Index of numbers is %d, and value of numbers is %d\n", index, value)
	}

	// here %c is for charater type variable and %d is for integer type variable.
	name := "Rahul Nagarwal"
	for index, char := range name {
		fmt.Printf("Index of name is %d, and value of name is %c\n", index, char)
	}
}
