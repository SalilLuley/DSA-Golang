package leetcode_75

import "fmt"

func IsSubsequence() {
	s := "b"
	t := "ahbgdc"
	fmt.Printf("isSubsequence(s, t): %v\n", isSubsequence(s, t))
}

// Time complexity: O(n)
// Space complexity: O(1)
func isSubsequence(s string, t string) bool {
	j := 0

	if s == t {
		return true
	}

	if len(s) == 0 {
		return true
	}

	for i := 0; i < len(t); i++ {
		if j < len(s) && t[i] == s[j] {
			j++
		}
	}
	return j == len(s)
}
