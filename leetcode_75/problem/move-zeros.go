package leetcode_75

import "fmt"

func MoveZeroes() {
	nums := []int{0, 1, 0, 3, 12}
	fmt.Printf("Move Zeros%d\n", moveZeroes(nums))
}

func moveZeroes(nums []int) []int {

	j := 0

	for i := range nums {
		if nums[i] != 0 {
			nums[j] = nums[i]
			j++
		}
	}

	for j < len(nums) {
		nums[j] = 0
		j++
	}

	return nums
}
