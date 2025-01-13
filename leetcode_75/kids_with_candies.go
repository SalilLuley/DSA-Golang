package leetcode_75

import "fmt"

func KidsWithCandies() {
	candies := []int{2, 3, 5, 1, 3}
	extraCandies := 3
	fmt.Printf("kidsWithCandies(): %v\n", kidsWithCandies(candies, extraCandies))
}

// Brute force approach
// Time Complexity: O(n)
// Space Complexity: O(n)
func kidsWithCandies(candies []int, extraCandies int) []bool {
	max := 0
	ans := []bool{}
	for _, v := range candies {
		if v > max {
			max = v
		}
	}
	for _, v := range candies {
		if v+extraCandies < max {
			ans = append(ans, false)
		} else {
			ans = append(ans, true)
		}
	}

	return ans
}
