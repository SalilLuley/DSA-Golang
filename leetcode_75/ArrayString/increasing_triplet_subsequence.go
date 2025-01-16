package leetcode_75

import "fmt"

// https://leetcode.com/problems/increasing-triplet-subsequence/description/?envType=study-plan-v2&envId=leetcode-75
func IncreasingTriplet() {
	nums := []int{4, 2, 1, 3, 5, 6}
	fmt.Printf("increasingTriplet(): %v\n", increasingTriplet(nums))
}

// Time complexity: O(n)
// Space complexity: O(1)
func increasingTriplet(nums []int) bool {
	if len(nums) < 3 {
		return false
	}

	first := int(^uint(0) >> 1)
	second := int(^uint(0) >> 1)

	for _, num := range nums {
		if num <= first {
			first = num
		} else if num <= second {
			second = num
		} else {
			return true
		}
	}

	return false
}
