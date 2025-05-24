package main

import (
	"errors"
	"fmt"
)

func Division(a, b float32) (float32, error) {
	if b == 0 {
		return 0, errors.New("0  is nit divicible by any number")
	}
	return a/b , nil
}

func main() {
	result , err:=Division(10,0)
	if err != nil{
		fmt.Println("Error :",err.Error())
		return
	}
	fmt.Println("The result is :" , result)
}