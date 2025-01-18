package leetcode_75

import (
	"fmt"
	"math"
)

func ContainerWithMostWater() {
	height := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	fmt.Printf("maxArea(height): %v\n", maxArea(height))
}

// Brute force
// Time complexity: O(n2)
// Space complexity: O(1)
// func maxArea(height []int) int {
// 	maxArea := 0
// 	for i := 0; i < len(height)-1; i++ {
// 		for j := i + 1; j < len(height); j++ {
// 			area := (j - i) * int(math.Min(float64(height[i]), float64(height[j])))
// 			if area > maxArea {
// 				maxArea = area
// 			}
// 		}
// 	}
// 	return maxArea
// }

// Time Complexty : O(n)
// Space Complexity: O(1)
func maxArea(height []int) int {
	l := 0
	r := len(height) - 1
	maxArea := 0
	for l < r {

		area := (r - l) * int(math.Min(float64(height[l]), float64(height[r])))
		if area > maxArea {
			maxArea = area
		}

		if height[r] > height[l] {
			l++
		} else {
			r--
		}

	}

	return maxArea
}
