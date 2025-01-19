package leetcode_75

import (
	"fmt"
	"math"
)

func FindHighestAltitude() {
	gain := []int{-5, 1, 5, 0, -7}
	fmt.Printf("largestAltitude(gain): %v\n", largestAltitude(gain))
}

// Time Complexity : O(n)
// Space Complexity: O(1)
func largestAltitude(gain []int) int {
	highestAltitude := 0
	sum := 0
	for i := 0; i < len(gain); i++ {
		sum += gain[i]
		highestAltitude = int(math.Max(float64(highestAltitude), float64(sum)))
	}

	return highestAltitude
}
