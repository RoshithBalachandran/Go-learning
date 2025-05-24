package main

import "fmt"

func divide(a int, b int) (int, int,int,int) {
	quotient := a / b
	remainder := a % b
	addition := a + b
	multiplication :=a * b
	return quotient, remainder, addition,multiplication
}

func main() {
	quotient, remainder, addition,multiplication := divide(10, 3)
	fmt.Println("Quotient:", quotient)
	fmt.Println("Remainder:", remainder)
	fmt.Println("Addition:", addition)
	fmt.Println("multiplication:", multiplication)
}
