package array

import "fmt"

func IncreasingTriplet() {
	nums := []int{5, 4, 3, 2, 1}
	fmt.Printf("increasingTriplet(nums) %v", increasingTriplet(nums))
}

// 5 - 2
//  2 3 4 5
// 0 1 2 3 4
// i j k
//   i j k
//     i j k

func increasingTriplet(nums []int) bool {

	if len(nums) < 3 {
		return false
	}

	for i := 0; i < len(nums)-2; i++ {
		if nums[i] < nums[i+1] && nums[i+1] < nums[i+2] {
			return true
		}
	}

	return false
}
