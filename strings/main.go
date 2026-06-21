package main

import (
	"fmt"
	"strings"
)

func main() {
	data := "Hello, World !"
	parts := strings.Split(data, ",")
	fmt.Println(parts)

	// Convert to uppercase
	// upperData := strings.ToUpper(data)
	// fmt.Println("Uppercase string:", upperData)

	// strings.Count
	// strings.TrimnSpace
	// strings.Replace

	str1 := "Rahul"
	str2 := "Nagarwal"
	concatenated := strings.Join([]string{str1, "Kumar", str2}, " ") // Joining str1 and str2 with a space in between
	fmt.Println("Concatenated string:", concatenated)
}
