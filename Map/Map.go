package main

import "fmt"

func main() {
	students := make(map[string]float64)
	students["Roshith"] = 89.5
	students["Renith"] = 93.5
	students["Rijith"] = 85.5

	students["Renith"] = 95.3             //To update and add elements to array
	students["Roshith"] = 90.5
	delete(students,"Roshith")              //to delete from array
	val,ok := students["Roshith"]        //To identify if it is in or not in
	fmt.Println(val,ok)
	fmt.Println("The student and their marks : ")
	for name,mark := range students{
	fmt.Printf("%s : %.2f \n" ,name,mark)
	}
}