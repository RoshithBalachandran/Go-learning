package main

import (
	"errors"
	"fmt"
)

func division(a, b int) (int, error) {
	if b == 0 {
		return -1, errors.New("Divided by Zero")
	}
	return a/b , nil
}

func main() {
	value,err := division(10 , 0)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("the result is =", value)
}

