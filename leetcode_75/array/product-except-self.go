package array

import "fmt"

func ProductExceptSelf() {
	nums := []int{1, 2, 3, 4}
	fmt.Printf("productExceptSelf(nums) %v\n", productExceptSelf(nums))
}

// [1 2 3 4]
// 2 * 3 * 4, 1 * 3 * 4, 1 * 2 * 4, 1 * 2 * 3
//[24,12,8,6]

func productExceptSelf(nums []int) []int {
	

	return nums
}
