package main

import "fmt"

func main() {
	students := make(map[string]float64)
	students["Roshith"] = 89.5
	students["Renith"] = 93.5
	students["Rijith"] = 85.5

	students["Renith"] = 95.3
	students["Roshith"] = 90.5

	fmt.Println("The student and their marks : ")
	for name,mark := range students{
	fmt.Printf("%s : %.2f \n" ,name,mark)
	}
}