package main

import (
	"Myapp/BasicArithametic/Arithametic"
	"fmt"
)

func main() {
	a := 10
	b := 4
	fmt.Println(Arithametic.Add(a, b))
	fmt.Println(Arithametic.Sub(a, b))
	fmt.Println(Arithametic.Mul(a, b))
	fmt.Println(Arithametic.Div(a, b))
}
