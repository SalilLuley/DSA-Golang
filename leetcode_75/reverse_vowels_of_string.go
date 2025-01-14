package leetcode_75

import (
	"fmt"
	"unicode"
)

// https://leetcode.com/problems/reverse-vowels-of-a-string/?envType=study-plan-v2&envId=leetcode-75
func ReverseVowels() {
	s := ".,"
	fmt.Printf("reverseVowels(s): %v\n", reverseVowels(s))
}

// Time complexity: O(n)
// Space complexity: O(1)

// The function `reverseVowels` has a time complexity of O(n), where n is the length of the input string `s`.
// This is because the function uses a two-pointer approach to traverse the string from both ends towards the center. Each character in the string is examined at most twice (once by each pointer), leading to a linear time complexity.

// The space complexity of the function is O(n) as well.
// This is due to the creation of the `arr` slice, which holds the characters of the input string. The space used by the `vowels` slice is constant, as it only contains a fixed number of vowel characters. However, since the size of `arr` is proportional to the size of the input string, the overall space complexity remains O(n).

// In summary, the time complexity is O(n) and the space complexity is O(n).

func isVowel(c rune) bool {
	c = unicode.ToLower(c)
	return c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u'
}
func reverseVowels(s string) string {
	i, j := 0, len(s)-1
	arr := []rune(s)
	for i < j {
		for i < j && !isVowel(arr[i]) {
			i++
		}

		for j > i && !isVowel(arr[j]) {
			j--
		}

		arr[i], arr[j] = arr[j], arr[i]
		i++
		j--
	}
	return string(arr)
}
