package main

import "fmt"

func main() {
	day := 30

	switch day {
	case 1:
		fmt.Println("Today is 1st day of month")
	case 2:
		fmt.Println("Today is 2nd day of month")
	case 3:
		fmt.Println("Today is 3rd day of month")
	default:
		fmt.Println("Unkown day of month.")
	}

	month := "January"

	switch month {
	case "January", "February", "March":
		fmt.Println("Winter")
	case "April", "May", "June":
		fmt.Println("Summer")
	default:
		fmt.Println("Other Season.")
	}
}
