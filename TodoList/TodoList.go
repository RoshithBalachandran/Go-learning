package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Student struct {
	name string
	age  int
}

var task []Student

func main() {
	fmt.Println("\n Welcome to Todo app")
	scanner := bufio.NewScanner(os.Stdin)
for{
	fmt.Println("\n Enter Your Choice")
	fmt.Println("1. Add Todo")
	fmt.Println("2. Read Todo")
	fmt.Println("3. Update Todo")
	fmt.Println("4. Delete Todo")
	fmt.Println("5. Exit Todo")
	scanner.Scan()
	Choice, err := strconv.Atoi(scanner.Text())
	if err != nil {
		fmt.Println("Invalid....!!")
	}
	switch Choice {
	case 1:
		Add(scanner)
	case 2:
		Read()
	case 3:
		Update(scanner)
	case 4:
		Delete(scanner)
	case 5:
		fmt.Println("Exit Sucessfully")
		return
	default:
		fmt.Println("Invalid Charactor")
	}
}}

func Add(scanner *bufio.Scanner) {
	fmt.Println("Enter your Name")
	scanner.Scan()
	input := scanner.Text()
	fmt.Println("Enter Your Age")
	scanner.Scan()
	Age, err := strconv.Atoi(scanner.Text())
	if err != nil {
		fmt.Printf("Invalid....!!!")
	}
	task = append(task, Student{input, Age})
	fmt.Println("Student Added Sucessfully")
}
func Read() {
		fmt.Println("Reading all the Students")
		if len(task) == 0 {
			fmt.Println("No Taks...!!")
		}
		for i, val := range task {
			fmt.Printf("%d. %s: %d\n", i+1, val.name,val.age)
		}
		fmt.Println("student listed Sucessfully")
	}
func Update(scanner *bufio.Scanner) {
	fmt.Println("Enter student index to update (starting from 1):")
	scanner.Scan()
	inp, err := strconv.Atoi(scanner.Text())
	if err != nil || inp < 1 || inp > len(task) {
		fmt.Println("Invalid index")
		return
	}

	index := inp - 1

	fmt.Println("Enter new name:")
	scanner.Scan()
	newName := scanner.Text()

	fmt.Println("Enter new age:")
	scanner.Scan()
	ageStr := scanner.Text()
	newAge, err := strconv.Atoi(ageStr)
	if err != nil {
		fmt.Println("Invalid age input")
		return
	}

	task[index].name = newName
	task[index].age = newAge

	fmt.Println("Student updated successfully.")
}
func Delete(scanner *bufio.Scanner){
	fmt.Println("Enter The number of student to delete")
	scanner.Scan()
	del,err:=strconv.Atoi(scanner.Text())
	if err != nil || del < 1 || del > len(task) {
		fmt.Println("Invalid index")
		return
	}
	task=append(task[:del-1],task[del:]...)
	fmt.Println("Task Updated sucessfully...!!")
}