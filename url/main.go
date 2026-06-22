package main

import (
	"fmt"
	"net/url"
)

func main() {
	// fmt.Println("from URL : Uniform Resource Locators")

	myURL := "https://www.example.com/path/to/resource?query=param#fragment"

	parsedURL, err := url.Parse(myURL)

	if err != nil {
		fmt.Println("error occured : ", err)
	}
	fmt.Println("Schema : ", parsedURL.Scheme)
	fmt.Println("Host : ", parsedURL.Host)
	fmt.Println("Path : ", parsedURL.Path)
	fmt.Println("RawQuerry : ", parsedURL.RawQuery)

}
