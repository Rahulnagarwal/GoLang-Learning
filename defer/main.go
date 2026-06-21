package main

import "fmt"

// 2 differ works as a LIFO.
func main() {
	fmt.Println("Start")
	defer fmt.Println("Middle") // defer will execute at the end of the function
	fmt.Println("End")
}
