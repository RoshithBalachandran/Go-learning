package main

import "fmt"

func main() {
	// age := 19
	// if age < 18 {
	// 	fmt.Println("You are a teenager")
	// }else if(age > 18){
	// 	fmt.Println("You are a adult")
	// }else{
	// 	fmt.Println("You are a 18")
	// }

	// const day = 8

	// switch day {
	// case 1:
	// 	fmt.Println("Munday")
	// 	case 2:
	// 	fmt.Println("Tuesday")
	// 	case 3:
	// 	fmt.Println("Wednesday")
	// 	case 4:
	// 	fmt.Println("Thursday")
	// 	case 5:
	// 	fmt.Println("Friday")
	// 	case 6:
	// 	fmt.Println("Saturday")
	// 	default :
	// 	fmt.Println("Sunday")
	// }
	// 	age:=15
	// if age >18{
	// 	fmt.Println("You are an adult" ,age)
	// }else{
	// 	fmt.Println("You are Teeneager" ,age)
	// }

	// switch time.Now().Weekday(){
	// case time.Saturday , time.Sunday:
	// 	fmt.Println("This is weekend",time.Now().Weekday())
	// default:
	// 	fmt.Println("This is work day",time.Now().Weekday())
	// }

	heii := func(i interface{}) {
		switch i.(type) {
		case int:
			fmt.Println("This is an integer")
		case string:
			fmt.Println("This is an String")
		case bool:
			fmt.Println("This is an Boolean")
		default:
			fmt.Println("Others")
		}
	}
	heii(3.14)
}
