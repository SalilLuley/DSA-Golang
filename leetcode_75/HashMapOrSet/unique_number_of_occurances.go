package leetcode_75

import (
	"fmt"
)

func UniqueNumberOfOccurances() {
	// arr := []int{1, 2, 2, 1, 1, 3}
	arr := []int{-3, 0, 1, -3, 1, 1, 1, -3, 10, 0}
	fmt.Printf("uniqueOccurrences(arr): %v\n", uniqueOccurrences(arr))
}

func uniqueOccurrences(arr []int) bool {
	hm := make(map[int]int)
	for i := 0; i < len(arr); i++ {
		value, ok := hm[arr[i]]
		if ok {
			value += 1
			hm[arr[i]] = value
		} else {
			hm[arr[i]] = 1
		}
	}

	set1 := make(map[int]bool)
	for i := 0; i < len(arr); i++ {
		if _, ok := set1[arr[i]]; !ok {
			set1[arr[i]] = true
		}
	}

	set2 := make(map[int]bool)
	for k := range set1 {
		if v, ok := hm[k]; ok {
			if _, okay := set2[v]; !okay {
				set2[v] = true
			}
		}
	}
	return len(set1) == len(set2)
}
