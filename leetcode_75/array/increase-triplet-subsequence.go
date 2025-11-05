package array

import (
	"fmt"
	"math"
)

func IncreasingTriplet() {
	nums := []int{1, 2, 1, 3}
	fmt.Printf("increasingTriplet(nums) %v", increasingTriplet(nums))
}

// fm < sm < num[i]
// 1 2 1 3
//
//       i
// fm >= i = 1
//  else fm <= sm && sm >= i = 1
//   else fm < sm < nums[i]

func increasingTriplet(nums []int) bool {
	fm := math.MaxInt
	sm := math.MaxInt

	for _, val := range nums {
		if fm >= val {
			fm = val
		} else if fm <= sm && sm >= val {
			sm = val
		} else if fm < sm && sm < val {
			return true
		}
	}
	return false
}
