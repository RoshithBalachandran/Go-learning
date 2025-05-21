package main

import "fmt"

func Calcualter(a, b int) (int, int, int, int) {
	sum := a + b
	difference := a - b
	product := a * b
	quotient := a / b
	return sum, difference, product, quotient
}
func main() {
	s,d,p,q:=Calcualter(10, 4)
	fmt.Println("Sum is ",s)
	fmt.Println("difference is ",d)
	fmt.Println("product is ",p)
	fmt.Println("quotient is ",q)
}