package main

import "fmt"

func main() {

	// using make keyword to create a map.
	// indialise by map word then key type and value type.
	studentGrades := make(map[string]int)

	studentGrades["a"] = 90
	studentGrades["b"] = 80
	studentGrades["c"] = 70

	// to delete.
	delete(studentGrades, "b")

	// checking if a key exists in the map
	grade, ok := studentGrades["a"]
	if ok {
		println("Grade of a is", grade)
	} else {
		println("Key a does not exist in the map")
	}

	// initialising a map using a map literal.
	user := map[string]int{
		"John": 500,
		"Jane": 600,
		"Bob":  550,
	}

	for name, salery := range user {
		fmt.Printf("---- Salery of %s is %d\n", name, salery)
	}
}
