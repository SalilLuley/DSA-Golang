package leetcode_75

import "fmt"

// https://leetcode.com/problems/product-of-array-except-self/description/?envType=study-plan-v2&envId=leetcode-75
func ProductExceptSelf() {
	nums := []int{1, 2, 3, 4}
	fmt.Printf("productExceptSelf(nums): %v\n", productExceptSelf(nums))
}

// Brute force approach
// Time complexity: O(n^2)
// func productExceptSelf(nums []int) []int {
// 	result := []int{}
// 	for i := 0; i < len(nums); i++ {
// 		prod := 1
// 		for j := 0; j < len(nums); j++ {
// 			if i != j {
// 				prod *= nums[j]
// 			}
// 		}
// 		result = append(result, prod)
// 	}

// 	return result
// }

// Time complexity: O(n)
// Space complexity: O(n)
func productExceptSelf(nums []int) []int {
	n := len(nums)
	left := make([]int, n)
	right := make([]int, n)

	left[0] = 1
	for i := 1; i < n; i++ {
		left[i] = left[i-1] * nums[i-1]
	}

	right[n-1] = 1

	for i := n - 2; i >= 0; i-- {
		right[i] = right[i+1] * nums[i+1]
	}

	for i := 0; i < n; i++ {
		nums[i] = left[i] * right[i]
	}

	return nums
}

// Best Solution

// func productExceptSelf(nums []int) []int {
// 	ans := make([]int, len(nums))
// 	ans[0] = 1

// 	for i := 1; i < len(nums); i++ {
// 		ans[i] = ans[i-1] * nums[i-1]
// 	}
// 	right := 1

// 	for j := len(nums) - 1; j >= 0; j-- {
// 		ans[j] = ans[j] * right
// 		right = right * nums[j]
// 	}
// 	return ans
// }
