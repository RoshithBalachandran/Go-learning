package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)
func main(){
	fmt.Println("Enter your birth year ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input ,_:= strconv.ParseInt(scanner.Text(),10,64)
	fmt.Printf("The age is %d before 2025 end",2025-input)

}