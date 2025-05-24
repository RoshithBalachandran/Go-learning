package main

import "fmt"

func mathametic(x, y int) (z ,z1 int) {
	defer fmt.Println("Check for it")
	z=x+y
	z1=x-y
	return
}

func main() {
	add , sub := mathametic(18, 8)
	fmt.Println(add , sub)

}