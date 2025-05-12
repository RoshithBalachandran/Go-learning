package main

import "fmt"

func main() {
	var n int
	fmt.Println("enter a number for Fibonacci")
	fmt.Scan(&n)
	a := 0
	b := 1
	var c int
	fmt.Print("Fibonacci series ")
	for i := 1; i <= n; i++ {
		c = a + b
		a = b
		b=c
		
		fmt.Print(a, " ")
	}
}
