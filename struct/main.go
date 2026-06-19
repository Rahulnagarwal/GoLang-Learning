package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

type Address struct {
	HouseNumber int
	Street      string
	City        string
	State       string
	ZipCode     string
}

// *** Type embedding
type Employee struct {
	Person_Details Person  // Embedded struct
	Person_Address Address // Embedded struct
	Salary         float64
}

func main() {
	// filling one by one.
	var p Person
	p.Name = "Alice"
	p.Age = 30

	fmt.Println("Name:", p.Name)
	fmt.Println("Age:", p.Age)

	// using struct literal
	user1 := Person{Name: "Bob", Age: 25}
	fmt.Println("Name:", user1)

	// new keyword -> it will allocate the memory, that's why in log we see the address(&) of the struct.
	user2 := new(Person)
	user2.Name = "Charlie"
	user2.Age = 35
	fmt.Println("Name:", user2)

	var Rahul Employee
	Rahul.Person_Details = Person{Name: "Nagarwal", Age: 28}
	Rahul.Person_Address = Address{HouseNumber: 123, Street: "Main St", City: "New York", State: "NY", ZipCode: "10001"}

	fmt.Println("Employee Name:", Rahul.Person_Details.Name)
	fmt.Println("Employee Zip Code:", Rahul.Person_Address.ZipCode)
}
