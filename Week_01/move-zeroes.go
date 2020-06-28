package main

import "fmt"

// [0,1,4,0,3,6]
func moveZeroes(nums []int) {
	var i, j int
	for i = 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[j], nums[i] = nums[i], nums[j]
			j++
		}
	}
}

func main() {
	nums := []int{0, 1, 4, 0, 3, 6}
	moveZeroes(nums)
	fmt.Println(nums)
}
