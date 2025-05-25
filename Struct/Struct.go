package main

import "fmt"

type Person struct {
	name string
	age  int
}

func (p Person) PrintDetails() {
	fmt.Println("Person name is ", p.name)
	fmt.Println("Person age is ", p.age)
}

func main() {

	person1 := Person{
		name: "Roshith",
		age:  24,
	}
	person1.PrintDetails()

}
