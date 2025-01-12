package leetcode_75

import "fmt"

func TwoSum() {
	// fmt.Println(twoSum())
	fmt.Printf("twoSumOpt(): %v\n", twoSumOpt())
}

// Brute force - Time Complexity - O(n2)
// Space Complexity - O(1)
// func twoSum() []int {
// 	arr := []int{6, 7, 1, 2, 3, 4, 5}
// 	target := 9
// 	for k, i := range arr {
// 		for n, j := range arr {
// 			if i+j == target {
// 				return []int{k, n}
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
func twoSumOpt() []int {
	arr := []int{6, 7, 1, 2, 3, 4, 5}
	target := 9
	hm := make(map[int]int)
	for i, v := range arr {
		numberToFind := target - v
		index, ok := hm[numberToFind]
		if ok {
			return []int{index, i}
		}
		hm[v] = i
	}

	return nil
}
