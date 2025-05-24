package main

import "fmt"

func valueChange(str *string) {
	*str = "Changed"
}

func main() {
	value := "Hellow Welcome !!!"
	fmt.Println(value)
	valueChange(&value)
	fmt.Println(value)
}