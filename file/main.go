package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

// WriteString writes the contents of the string s to w, which accepts a slice of bytes.
// os.Open,
// make([]byte, 1024) creates a buffer of size 1024 bytes to hold the file contents.
// file.Read(buffer) reads and store content into buffer.
// fmt.Print(string(buffer[:n])) ==> prints the contents read from the file, converting the byte slice to a string.
// ioutil.ReadFile can cause issue bcz it loads whole file into memory, which can be problematic for large files.

func main() {
	// file, err := os.Create("ByCode.txt")

	// if err != nil {
	// 	fmt.Println("Error creating file:", err)
	// }
	// defer file.Close()

	// content := "Hello World!"
	// _, errors := io.WriteString(file, content)

	// if errors != nil {
	// 	fmt.Println("Error writing to file:", errors)
	// 	return
	// }
	// fmt.Println("File created and content written successfully.")

	file, err := os.Open("ByCode.txt")

	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	// Create a buffer to hold the file contents
	buffer := make([]byte, 1024)

	// Read the file contents into the buffer
	for {
		n, err := file.Read(buffer)
		if err == io.EOF {
			break // End of file reached
		}
		if err != nil {
			fmt.Println("Error reading file:", err)
			return
		}

		// Print the contents read from the file
		fmt.Print(string(buffer[:n]))
	}

	// read the file contents again via os.ReadFile
	content, err := ioutil.ReadFile("ByCode.txt")
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Print the contents read from the file
	fmt.Println(string(content))
}
