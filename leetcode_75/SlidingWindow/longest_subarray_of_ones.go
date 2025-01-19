package leetcode_75

import (
	"fmt"
	"math"
)

func LongestSubarrayOfOnes() {
	nums := []int{1, 1, 0, 1}
	fmt.Printf("longestSubarray(nums): %v\n", longestSubarray(nums))
}

func longestSubarray(nums []int) int {
	start := 0
	index := -1
	maxSubArray := 0

	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			start = index + 1
			index = i
		}
		maxSubArray = int(math.Max(float64(maxSubArray), float64(i)-float64(start)))
	}

	return maxSubArray
}
