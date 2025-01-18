package leetcode_75

import "fmt"

func MaxNumberOfKSumPairs() {
	nums := []int{2, 5, 4, 4, 1, 3, 4, 4, 1, 4, 4, 1, 2, 1, 2, 2, 3, 2, 4, 2}

	// nums := []int{1, 2, 3, 4}
	k := 3
	fmt.Printf("maxOperations(nums, k): %v\n", maxOperations(nums, k))
}

// Brute force
//
//	func maxOperations(nums []int, k int) int {
//		counter := 0
//		for l := 0; l < len(nums)-1; l++ {
//			for r := l + 1; r < len(nums); r++ {
//				if nums[l]+nums[r] == k {
//					nums[l] = 0
//					nums[r] = 0
//					counter++
//				}
//			}
//		}
//		return counter
//	}
//
// 3,1,3,4,3
func maxOperations(nums []int, k int) int {
	hm := make(map[int]int, 1)
	pairs := 0
	for i := 0; i < len(nums); i++ {
		compensate := k - nums[i]
		if value, ok := hm[compensate]; ok && value >= 1 {
			value -= 1
			hm[compensate] = value
			pairs += 1
		} else if value, ok := hm[nums[i]]; ok {
			value += 1
			hm[nums[i]] = value
		} else {
			hm[nums[i]] = 1
		}
	}
	return pairs
}
