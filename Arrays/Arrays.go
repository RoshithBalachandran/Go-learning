package main

import "fmt"

func main() {
	var nums = [5]int{8, 5, 3, 4, 1}

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
	fmt.Println("The minimum number is :", min)
	fmt.Println("The minimum number is :", max)
}
