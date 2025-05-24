package main

import "fmt"

func main() {
	day := 4

	switch day {
	case 1:
		fmt.Printf("Monday")
	case 2:
		fmt.Printf("Tuesday")
	case 3:
		fmt.Printf("Wednesday")
	case 4:
		fmt.Printf("Thursday")
	case 5:
		fmt.Printf("Friday")
	case 6:
		fmt.Printf("Saturday")
		default :
		fmt.Print("I Dont Remember")
	}
}
