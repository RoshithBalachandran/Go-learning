package main

import "fmt"

func main() {
	var nums = [8]int{5, 9, 7, 6, 8, 2, 4, 1}
	min := nums[0]
	max := nums[0]

	for i := 0; i < len(nums); i++ {
		if nums[i] < min {
			min = nums[i]
		}
		if nums[i] > max {
			max = nums[i]
		}
	}
	fmt.Println("The min is :",min)
	fmt.Println("the max is :",max)
}

