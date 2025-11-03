package array

import (
	"fmt"
	"strings"
)

func ReverseVowels() {
	s := "IceCreAm"
	fmt.Printf("reverseVowels(), %v\n", reverseVowels(s))
}

func reverseVowels(s string) string {
	vowels := "AEIOUaeiou"
	runes := []rune(s)
	start := 0
	end := len(s) - 1

	for start < end {
		for start < end && !strings.Contains(vowels, string(runes[start])) {
			start++
		}

		for end > start && !strings.Contains(vowels, string(runes[end])) {
			end--
		}

		if start < end {
			runes[start], runes[end] = runes[end], runes[start]
			start++
			end--
		}
	}

	return string(runes)
}
