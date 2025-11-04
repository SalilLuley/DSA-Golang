package array

import "fmt"

// [1 2 3 4]
// 2 * 3 * 4, 1 * 3 * 4, 1 * 2 * 4, 1 * 2 * 3
//[24,12,8,6]

func ProductExceptSelf() {
	nums := []int{-1, 1, 0, -3, 3}
	fmt.Printf("productExceptSelf() %v", productExceptSelf(nums))
}

// func productExceptSelf(nums []int) []int {
// 	sum := 1
// 	results := []int{}
// 	for _, val := range nums {
// 		sum *= val
// 	}
// 	for _, val := range nums {
// 		results = append(results, sum/val)
// 	}
// 	return results
// }

func productExceptSelf(nums []int) []int {
	res := []int{}
	prefix := 1
	for _, val := range nums {
		res = append(res, prefix)
		prefix *= val
	}

	suffix := 1

	for i := len(nums) - 1; i >= 0; i-- {
		res[i] *= suffix
		suffix *= nums[i]
	}

	return res
}
