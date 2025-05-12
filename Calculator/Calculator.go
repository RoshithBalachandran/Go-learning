package main

import "fmt"

func main() {
	var a int
	var b int
	var c int
	fmt.Println("enter first number : ")
	fmt.Scan((&a))

	fmt.Println("enter Second number : ")
	fmt.Scan((&b))

	fmt.Println("1 for addition \n2 for subtraction \n3 for multiplicaton \n4 for division ")
	fmt.Scan(&c)

	switch c {
	case 1:
		fmt.Println(a + b)
	case 2:
		fmt.Println(a - b)
	case 3:
		fmt.Println(a * b)
	case 4:
		fmt.Println(a / b)
	}

}
