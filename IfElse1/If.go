package main

import "fmt"

func main() {
	var num1 = 2
	var num2 = 5
	var num3 = 10
	if num1 > num2 && num1 > num3 {
		fmt.Println("num1 is greater than num2 and num3")
	}else if num2>num1&& num2>num3{
			fmt.Println("num2 is greater than num1 and num3")
	}else{
			fmt.Println("num3 is greater than num2 and num1")
	}
}