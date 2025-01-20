package leetcode_75

import (
	"fmt"
	"sort"
)

func DetermineIfTwoStringsAreClose() {
	word1 := "aacabb"
	word2 := "bbcbaa"
	fmt.Printf("closeStrings(word1, word2): %v\n", closeStrings(word1, word2))
}

// Time Complexity O (n log n)
// Space Complexity O (n)
func closeStrings(word1 string, word2 string) bool {
	if len(word2) != len(word1) {
		return false
	}
	hm1 := make(map[byte]int)
	hm2 := make(map[byte]int)

	for i := 0; i < len(word1); i++ {
		hm1[word1[i]] = hm1[word1[i]] + 1
	}

	for i := 0; i < len(word2); i++ {
		hm2[word2[i]] = hm2[word2[i]] + 1
	}

	arr1 := []int{}
	for k := range hm1 {
		if _, ok := hm2[k]; ok {
			arr1 = append(arr1, hm1[k])
		}
	}

	arr2 := []int{}
	for k := range hm2 {
		arr2 = append(arr2, hm2[k])
	}
	sort.Ints(arr1)
	sort.Ints(arr2)
	return equal(arr1, arr2)
}

func equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
