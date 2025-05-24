package main

import "fmt"

func main()  {
	var a int
	fmt.Println("Enter a number")
	fmt.Scan(&a)

	if a < 0 {
		fmt.Println(" number is negative")
	}else if a>0{
		fmt.Println(" number is Possitive")
	}else{
		fmt.Println(" number is Zerow")
	}
	
}