package main

import "fmt"

func RemoveDooplicate(arr []int) []int {
	unique := []int{}
	for i := 0; i < len(arr); i++ {
		found := false
		for j := 0; j < len(unique); j++ {
			if arr[i] == arr[j] {
				found = true
				break
			}
		}
		if !found {
			unique = append(unique, arr[i])
		}
	}
	return unique
}

func main() {
	array := []int{2, 9, 7, 6, 4, 3, 1, 2, 4, 9, 5}
	result := RemoveDooplicate(array)
	fmt.Println("the array is : ",array)
	fmt.Println("After removing Dooplicate : ",result)
}

