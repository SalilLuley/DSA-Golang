package leetcode_75

import (
	"fmt"
)

func MaximumAverageSubarray() {
	nums := []int{1, 12, -5, -6, 50, 3}
	k := 4
	fmt.Printf("findMaxAverage(nums, k): %v\n", findMaxAverage(nums, k))
}

// Time Complexity : O(n)
// Space Complexity : O(1)
//
//	func findMaxAverage(nums []int, k int) float64 {
//		if len(nums) == 0 {
//			return 0
//		}
//
//		if len(nums) == 1 {
//			return float64(nums[0])
//		}
//		i := 0
//		j := 1
//		average := 0.00000
//		sum := nums[0]
//		for i < len(nums)-k {
//			fmt.Printf("i: %v\n", i)
//			for j < k {
//				sum += nums[i+j]
//				j++
//			}
//			avg := float64(sum) / float64(k)
//			if average < avg {
//				average = avg
//			}
//			i++
//			j = 1
//			sum = nums[i]
//		}
//		return average
//	}
//
// Time Complexity O(n*k)
// Space Complexity O(1)
func findMaxAverage(nums []int, k int) float64 {
	if len(nums) == 0 {
		return 0
	}

	if len(nums) == 1 {
		return float64(nums[0])
	}

	maxSum := -1.0 * float64(^uint(0)>>1) // Initialize to the smallest possible float64 value
	for i := 0; i <= len(nums)-k; i++ {
		sum := 0
		for j := 0; j < k; j++ {
			sum += nums[i+j]
		}
		if float64(sum) > maxSum {
			maxSum = float64(sum)
		}
	}
	return maxSum / float64(k)
}
