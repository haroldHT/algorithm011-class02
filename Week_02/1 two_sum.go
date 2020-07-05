package main

//func twoSum(nums []int, target int) []int {
//	m := make(map[int]int)
//	var pre int
//	var ok bool
//
//	for i := 0; i < len(nums); i++ {
//		if pre, ok = m[target-nums[i]]; ok {
//			return []int{pre, i}
//		} else {
//			m[nums[i]] = i
//		}
//	}
//
//	return []int{}
//}

func twoSum(nums []int, target int) []int {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}
	return []int{}
}
