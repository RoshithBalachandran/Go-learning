package main

import "fmt"

func main() {
	n := 6
	for i := 0; i <= n; i++ {
		for j := 0; j <= n; j++ {
			if i == 0 || j == 0 || j == i || j == n || i == n || j == n-i {
				fmt.Print("* ")
			} else {
				fmt.Print("  ")
			}
		}
		fmt.Println()
	}

}
