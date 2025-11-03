package array

import (
	"fmt"
	"strings"
)

func ReverseWords() {
	s := "the sky is  blue"
	fmt.Printf("reverseWords() %v \n", reverseWords(s))
}

// "a good   example" []
// 0 1 2 3  4 5

// "example good a"

func reverseWords(s string) string {
	arr := strings.Split(s, " ")
	result := []string{}
	lenArr := len(arr) - 1
	for i := lenArr; i >= 0; i-- {
		if len(arr[i]) > 0 {
			result = append(result, arr[i])
		}
	}
	return strings.Join(result, " ")
}
