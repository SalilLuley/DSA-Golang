package leetcode_75

import (
	"fmt"
)

func MaximumAverageSubarray() {
	// nums := []int{1, 12, -5, -6, 50, 3}
	nums := []int{0, 4, 0, 3, 2}
	k := 1
	fmt.Printf("findMaxAverage(nums, k): %v\n", findMaxAverage(nums, k))
}

// Time Complexity : O(n2)
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
//
//	func findMaxAverage(nums []int, k int) float64 {
//		if len(nums) == 0 {
//			return 0
//		}
//
//		if len(nums) == 1 {
//			return float64(nums[0])
//		}
//
//		maxSum := -1.0 * float64(^uint(0)>>1) // Initialize to the smallest possible float64 value
//		for i := 0; i <= len(nums)-k; i++ {
//			sum := 0
//			for j := 0; j < k; j++ {
//				sum += nums[i+j]
//			}
//			if float64(sum) > maxSum {
//				maxSum = float64(sum)
//			}
//		}
//		return maxSum / float64(k)
//	}
//
// Time Complexity  : O(n*m)
// Space Complexity : O(1)
func findMaxAverage(nums []int, k int) float64 {
	lenArr := len(nums)
	if lenArr == 0 {
		return 0
	}

	sum := 0

	for i := 0; i < k; i++ {
		sum += nums[i]
	}
	maxSum := sum
	for i := k; i < len(nums); i++ {
		sum = sum + nums[i] - nums[i-k]
		if sum > maxSum {
			maxSum = sum
		}
	}

	return float64(maxSum) / float64(k)

}
