package main

import "fmt"

func main() {
	var num int
	fmt.Println("enter a number")
	fmt.Scan(&num)
	if(num%2==0){
		fmt.Println("Even Numbers")
	}else{
		fmt.Println("Odd numbers")
	}
}