package main

import "fmt"

// Syntax : numbers := []int{1, 2, 3}
// Capacity doubles every time whenever it reaches the lenght : 5 -> 10 -> 20 -> 40
func main() {
	// numbers := []int{1, 2, 3, 4, 5}
	// numbers = append(numbers, 6, 7, 8)
	// fmt.Println("Number :", numbers)
	// fmt.Printf("Number has data type : %T\n", numbers)
	// fmt.Println("Length :", len(numbers))

	numbers := make([]int, 3, 5)

	numbers = append(numbers, 4)
	numbers = append(numbers, 98)
	numbers = append(numbers, 5)

	fmt.Println("Slice :", numbers)
	fmt.Println("Length :", len(numbers))
	fmt.Println("Capacity :", cap(numbers))
}
