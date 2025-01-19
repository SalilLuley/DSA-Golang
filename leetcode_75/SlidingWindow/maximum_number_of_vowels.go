package leetcode_75

import (
	"fmt"
)

func MaximumNumberOfVowelsInASubstringOfGivenLength() {
	s := "abciiidef"
	k := 3
	fmt.Printf("maxVowels: %v\n", maxVowels(s, k))
}

// func maxVowels(s string, k int) int {
// 	countVowels := 0
// 	maxVowels := 0
// 	innerLoop := 0
// 	str := "aeiouAEIOU"
// 	for i := 0; i <= len(s)-k; i++ {
// 		for j := i; innerLoop < k; j++ {
// 			if strings.Contains(str, string(s[j])) {
// 				countVowels++
// 			}
// 			innerLoop++
// 		}
// 		if maxVowels < countVowels {
// 			maxVowels = countVowels
// 		}
// 		innerLoop = 0
// 		countVowels = 0
// 	}
// 	return maxVowels
// }

func maxVowels(s string, k int) int {
	isVowel := func(c byte) bool {
		return c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u' ||
			c == 'A' || c == 'E' || c == 'I' || c == 'O' || c == 'U'
	}

	countVowels := 0
	maxVowels := 0

	// Initialize the first window
	for i := 0; i < k; i++ {
		if isVowel(s[i]) {
			countVowels++
		}
	}
	maxVowels = countVowels

	// Slide the window across the string
	for i := k; i < len(s); i++ {
		if isVowel(s[i]) {
			countVowels++
		}
		if isVowel(s[i-k]) {
			countVowels--
		}
		if countVowels > maxVowels {
			maxVowels = countVowels
		}
	}

	return maxVowels
}
