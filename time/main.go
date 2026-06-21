package main

import (
	"fmt"
	"time"
)

func main() {

	currentTime := time.Now()
	fmt.Println("Current time:", currentTime)
	fmt.Printf("Type %T\n", currentTime)

	// ** here fomatting is little bit different than normal languages.
	// 2006-01-02 15:04:05 Monday is the reference time in Go.
	formattedTime := currentTime.Format("02-01-2006, 15:04:05 Monday")
	fmt.Println("Formatted time:", formattedTime)

	// Reference time for parsing
	a := "2006-01-02" // This is the reference time format for parsing
	b := "2023-06-15"
	parsedTime, _ := time.Parse(a, b)
	fmt.Println("Parsed time:", parsedTime) // Parsed time: 2023-06-15 00:00:00 +0000 UTC

	// add 2 more days
	added_date := currentTime.Add(48 * time.Hour)
	fmt.Println("Added date:", added_date)
	formatted_added_date := added_date.Format("2006/02/01, Monday")
	fmt.Println("Formatted added date:", formatted_added_date)
}
