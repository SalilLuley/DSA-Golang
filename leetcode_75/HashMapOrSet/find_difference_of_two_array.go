package leetcode_75

import "fmt"

func FindDifferenceOfTwoArrays() {
	nums1 := []int{1, 2, 3, 3}
	nums2 := []int{1, 1, 2, 2}

	fmt.Printf("findDifference(nums1, nums2): %v\n", findDifference(nums1, nums2))
}

// Time Complexity O ( m+n )
// Space Complexity O ( m+n )
func findDifference(nums1 []int, nums2 []int) [][]int {
	set1 := make(map[int]bool)
	set2 := make(map[int]bool)

	for _, num := range nums1 {
		set1[num] = true
	}
	for _, num := range nums2 {
		set2[num] = true
	}

	diff1 := []int{}
	for num := range set1 {
		if !set2[num] {
			diff1 = append(diff1, num)
		}
	}

	diff2 := []int{}
	for num := range set2 {
		if !set1[num] {
			diff2 = append(diff2, num)
		}
	}

	return [][]int{diff1, diff2}
}
