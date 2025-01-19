package leetcode_75

import (
	"fmt"
	"math"
)

func MaxConsecutiveOnes() {
	nums := []int{0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 1, 1, 0, 0, 0, 1, 1, 1, 1}
	k := 3
	fmt.Printf("longestOnes(nums, k): %v\n", longestOnes(nums, k))
}

// Time Complexity = O(n)
// Space Complexity = O(1)
func longestOnes(nums []int, k int) int {
	start := 0
	end := 0
	maxZeros := 0
	for end < len(nums) {
		if nums[end] == 0 {
			k -= 1
		}
		for k < 0 {
			if nums[start] == 0 {
				k += 1
			}
			start += 1
		}
		maxZeros = int(math.Max(float64(maxZeros), float64(end)-float64(start)+1))
		end++
	}
	return maxZeros
}
