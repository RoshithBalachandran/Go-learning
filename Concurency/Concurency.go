package main

import (
	"fmt"
	"time"
)

func printOdd() {
	for i := 1; i <= 10; i += 2 {
		fmt.Println("Odd", i)
		time.Sleep(1 * time.Second)
	}
}
func printEven() {
	for i := 0; i <= 10; i += 2 {
		fmt.Println("Even", i)
		time.Sleep(1 * time.Second)
	}
}

func main() {

	go printOdd()

	go printEven()
	time.Sleep(3 * time.Second)
}
