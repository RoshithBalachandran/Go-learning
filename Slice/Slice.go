package main

import "fmt"

func main() {
	var nums = []int{1, 2, 3}
	nums = append(nums, 4, 5)
	fmt.Println(nums , len(nums))
	var NewNums=make([]int ,len(nums))
	copy(NewNums , nums)
	NewNums =append(NewNums, 6,7,8,9)
	fmt.Println(NewNums,len(NewNums))
}