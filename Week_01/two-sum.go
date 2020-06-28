package main

import "fmt"

// [0,1,4,0,3,6] target 6
func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	var pre int
	var ok bool

	for i := 0; i < len(nums); i++ {
		if pre, ok = m[target-nums[i]]; ok {
			return []int{pre, i}
		} else {
			m[nums[i]] = i
		}
	}

	return []int{}
}

func main() {
	nums := []int{0, 1, 4, 0, 3, 6}

	fmt.Println(twoSum(nums, 6))
}
