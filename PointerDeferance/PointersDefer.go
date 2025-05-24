package main

import "fmt"

func changeValue(val *int) {
	*val = 20
}

func main() {
	var x = 10
	fmt.Println("Value : ",x)
	changeValue(&x)
	fmt.Println("Value after modifing using pointer : ",x)
}