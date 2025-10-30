package leetcode_75

import "fmt"

func TwoSum() {
	nums := []int{2, 7, 11, 15}
	target := 9
	fmt.Printf("TwoSum() %v \n", twoSum(nums, target))
}

func twoSum(nums []int, target int) []int {
	numsMap := make(map[int]int)
	for i, num := range nums {
		if val, ok := numsMap[target-num]; ok {
			return []int{i, val}
		}
		numsMap[num] = i
	}
	return nil
}
