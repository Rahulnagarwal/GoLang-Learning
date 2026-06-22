package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	fmt.Println("Web service")

	res, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		fmt.Println("Error making GET request:", err)
		return
	}
	defer res.Body.Close()

	fmt.Println("Response :", res)

	// Read the response body
	data, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	// *** Pass the data into string.
	fmt.Println("Response Body:", string(data))
}
