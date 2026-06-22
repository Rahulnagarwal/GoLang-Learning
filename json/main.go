package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	// for defining the exporting key should start with capital letter.
	Name    string `json:"name"`
	Age     int    `json:"age"`
	IsAdult bool   `json:"isAdult"`
}

func main() {
	rahul := Person{Name: "Rahul", Age: 30, IsAdult: true}
	fmt.Println("Person", rahul) // {Rahul 30 true}

	// converting into json encoding (marshelling)
	jsonData, err := json.Marshal(rahul)
	if err != nil {
		fmt.Println("Error encoding JSON:", err)
		return
	}
	fmt.Println("JSON Data:", string(jsonData)) // {"name":"Rahul","age":30,"isAdult":true}

	// converting back to struct (unmarshalling)
	var personFromJSON Person
	err = json.Unmarshal(jsonData, &personFromJSON)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}
	fmt.Println("Person from JSON:", personFromJSON) // {Rahul 30 true}
}
