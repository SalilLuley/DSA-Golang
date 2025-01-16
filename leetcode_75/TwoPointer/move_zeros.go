package leetcode_75

import "fmt"

func MoveZeros() {
	nums := []int{1, 0, 1, 0, 3, 12}
	moveZeros(nums)
}

// Time complexity: O(n)
// Space complexity: O(1)
func moveZeros(nums []int) {
	if len(nums) <= 1 {
		return
	}

	nonZeroIndex := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[nonZeroIndex], nums[i] = nums[i], nums[nonZeroIndex]
			nonZeroIndex++
		}
	}

	fmt.Printf("moveZeros(): %v\n", nums)
}
