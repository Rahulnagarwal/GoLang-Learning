package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Enter your name:")
	//var name string

	// Take input till first whitespace.
	// for whole string buf.io package can be used.
	// fmt.Scan(&name)
	// fmt.Println("Hello, Mr.", name)

	reader := bufio.NewReader(os.Stdin)
	name, _ := reader.ReadString('\n') // read till newline character
	fmt.Println("Hello, Mr.", name)
}
