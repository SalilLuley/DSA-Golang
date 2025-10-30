package array

import (
	"fmt"
)

func MergeAlternately() {
	word1 := "abc"
	word2 := "pqrst"
	fmt.Printf("mergeAlternately() %v\n", mergeAlternately(word1, word2))
}

func mergeAlternately(word1 string, word2 string) string {
	temp := ""
	maxNum := max(len(word1), len(word2))
	for i := 0; i < maxNum; i++ {
		if i < len(word1) {
			temp += string(word1[i])
		}

		if i < len(word2) {
			temp += string(word2[i])
		}
	}
	return temp
}
