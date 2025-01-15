package leetcode_75

import "fmt"

// https://leetcode.com/problems/two-sum/description/
func TwoSum() {
	fmt.Printf("twoSumOpt(): %v\n", twoSum())
}

// Brute force - Time Complexity - O(n2)
// Space Complexity - O(1)
// func twoSumBruteForce() []int {
// 	arr := []int{2, 7, 11, 15}
// 	target := 9
// 	for index1, i := range arr {
// 		for index2, j := range arr {
// 			if i+j == target {
// 				return []int{index1, index2}
// 			}
// 		}
// 	}
// 	return nil
// }

// Optimized approach
/**
Time Complexity - O(n)
Space Complexity - O (n)
*/
func twoSum() []int {
	arr := []int{2, 7, 11, 15}
	target := 9
	hm := make(map[int]int)
	for index, value := range arr {
		numberToFind := target - value
		if indexFound, ok := hm[numberToFind]; ok {
			return []int{indexFound, index}
		}
		hm[value] = index
	}
	return nil
}
