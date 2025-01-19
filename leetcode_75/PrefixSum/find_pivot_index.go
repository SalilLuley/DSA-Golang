package leetcode_75

import "fmt"

func FindPivotIndex() {
	nums := []int{1, 2, 3}
	fmt.Printf("pivotIndex(nums): %v\n", pivotIndex(nums))
}

//	func pivotIndex(nums []int) int {
//		start := 0
//		end := len(nums) - 1
//		leftSum := 0
//		rightSum := 0
//		mid := (start + (end - start) + 1) / 2
//
//		for mid >= 0 && mid < len(nums) {
//			fmt.Printf("mid %v\n", mid)
//			for start < mid {
//				leftSum += nums[start]
//				start++
//			}
//			fmt.Printf("leftSum: %v\n", leftSum)
//			for end > mid {
//				fmt.Printf("nums[end]: %v\n", nums[end])
//				rightSum += nums[end]
//				end--
//			}
//			fmt.Printf("rightSum: %v\n", rightSum)
//			if leftSum == rightSum {
//				return mid
//			} else if leftSum > rightSum {
//				mid--
//			} else {
//				mid++
//			}
//			start = 0
//			end = len(nums) - 1
//			leftSum = 0
//			rightSum = 0
//		}
//		return -1
//	}
//
// Time Complexity O(n)
// Space Complexity O(1)
func pivotIndex(nums []int) int {
	totalSum := 0
	for _, num := range nums {
		totalSum += num
	}

	leftSum := 0
	for i, num := range nums {
		if leftSum == totalSum-leftSum-num {
			return i
		}
		leftSum += num
	}

	return -1
}
