package leetcode_75

import (
	"fmt"
	"strings"
)

// https://leetcode.com/problems/reverse-words-in-a-string/description/?envType=study-plan-v2&envId=leetcode-75
func ReverseWords() {
	s := "  hello world  "
	fmt.Printf("reverseVowelsOfString(): %v\n", reverseWords(s))
}

// Time complexity: O(n)
// Space complexity: O(n)
func reverseWords(s string) string {
	arr := strings.Split(s, " ")
	result := make([]string, 0)
	for i := len(arr) - 1; i >= 0; i-- {
		if len(arr[i]) > 0 {
			result = append(result, arr[i])
		}
	}
	return strings.Join(result, " ")
}
