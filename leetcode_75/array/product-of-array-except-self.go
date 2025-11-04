package array

import (
	"fmt"
)

// [1 2 3 4]
// 2 * 3 * 4, 1 * 3 * 4, 1 * 2 * 4, 1 * 2 * 3
//[24,12,8,6]

func ProductExceptSelf() {
	nums := []int{-1, 1, 0, -3, 3}
	fmt.Printf("productExceptSelf() %v\n", productExceptSelf(nums))
}

// Does not work with 0
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

// Left
// prefix - 1
// res - [1, 1, 2, 6]
// 1*1, 1*2 = 2 , 2 * 3 = 6, 6 * 6 = 36

// right
// nums[1,2,3,4]
// res[1,1,2,6]
// suffix - 1/4/12/24
// res[i] = 6*1, 2*4, 12*1, 24*1
// suffix = 1*4, 4*3, 12*2,
// pos - 3

// -1,1,0,-3,3
// prefix - 1/-1/-1/0/0/0
// res = [1, -1, -1, 0, 0, 0]
// prefix = 1*-1, -1*1, -1*0, 0*-3,0*3

//suffix - 1
//res = i ==

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
