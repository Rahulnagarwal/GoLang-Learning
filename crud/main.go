package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

type Todo struct {
	UserID    int    `json:"userId"`
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func getTodo() {
	res, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		fmt.Println("Error fetching data:", err)
		return
	}
	defer res.Body.Close()

	// data, err := io.ReadAll(res.Body)
	// if res.StatusCode != http.StatusOK {
	// 	fmt.Println("Status is not OK", res.StatusCode)
	// 	return
	// }
	// if err != nil {
	// 	fmt.Println("Error reading response body:", err)
	// 	return
	// }

	// fmt.Println("Response Data:", string(data))

	var todo Todo
	// NewDecoder returns a new decoder that reads from r.
	err = json.NewDecoder(res.Body).Decode(&todo)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}

	fmt.Println("Todo:", todo)
}

// struct to json -> json.Marshal()
// json to string -> string(jsonData)
// string to reader -> strings.NewReader(jsonString)

func postTodo() {
	todo := Todo{
		UserID:    1,
		ID:        1,
		Title:     "Rahul Nagarwal",
		Completed: false,
	}

	// convert string to json
	jsonData, err := json.Marshal(todo)
	if err != nil {
		fmt.Println("Error encoding JSON:", err)
		return
	}

	// convert json to string
	jsonString := string(jsonData)

	// convert string to reader
	jsonReader := strings.NewReader(jsonString)
	myURL := "https://jsonplaceholder.typicode.com/todos"

	// third parameter of http.Post is of type io.Reader.
	res, err := http.Post(myURL, "application/json", jsonReader)
	if err != nil {
		fmt.Println("Error making POST request:", err)
		return
	}
	defer res.Body.Close()

	fmt.Println("Response Status:", res.Status)
}

// struct to json -> json.Marshal()
// json to string -> string(jsonData)
// string to reader -> strings.NewReader(jsonString)

// making a request using http.NewRequest().
// sending it using http.Client.Do()
func updateTodo() {
	todo := Todo{
		UserID:    34324,
		Title:     "Rahul Nagarwal",
		Completed: true,
	}

	// convert struct to json
	jsonData, err := json.Marshal(todo)
	if err != nil {
		fmt.Println("Error encoding JSON:", err)
		return
	}

	// convert json to string
	jsonString := string(jsonData)

	// convert string to reader
	jsonReader := strings.NewReader(jsonString)

	myURL := "https://jsonplaceholder.typicode.com/todos/1"

	// third parameter of http.Post is of type io.Reader.
	req, err := http.NewRequest("PUT", myURL, jsonReader)
	if err != nil {
		fmt.Println("Error making POST request:", err)
		return
	}
	req.Header.Set("Content-Type", "application/json")

	// send the request
	clien := http.Client{}
	res, err := clien.Do(req)
	if err != nil {
		fmt.Println("Error making PUT request:", err)
		return
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	fmt.Println("Response Status:", string(data))
}

func deleteTodo() {

	myURL := "https://jsonplaceholder.typicode.com/todos/1"

	req, err := http.NewRequest("DELETE", myURL, nil)
	if err != nil {
		fmt.Println("Error making DELETE request:", err)
		return
	}

	// send the request
	client := http.Client{}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("Error making DELETE request:", err)
		return
	}
	defer res.Body.Close()

	fmt.Println("Response Status:", res.Status)
}
func main() {
	// getTodo()
	// postTodo()
	// updateTodo()
	deleteTodo()
}
