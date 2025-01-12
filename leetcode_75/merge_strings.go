package leetcode_75

import (
	"fmt"
	"math"
	"strings"
)

func MergeStrings() {
	fmt.Printf("mergeStrings(): %v\n", mergeStrings())
}

// Time complexity = O(m+n)
// Space complexity = O(m+n)
// As modifying a string is costly operation we use array to do it.
func mergeStrings() string {
	word1 := "ab"
	word2 := "pqrstuv"
	arr1 := strings.Split(word1, "")
	arr2 := strings.Split(word2, "")
	str3 := []string{}
	counter1 := 0
	counter2 := 0
	min := int(math.Min(float64(len(arr1)), float64(len(arr2))))
	for i := 0; i < len(arr1)+len(arr2); i++ {
		if i < min*2 {
			if i%2 == 0 {
				str3 = append(str3, arr1[counter1])
				counter1++
			} else {
				str3 = append(str3, arr2[counter2])
				counter2++
			}
		} else {
			if i < len(arr1)*2 {
				str3 = append(str3, arr1[counter1])
				counter1++
			} else if i < len(arr2)*2 {
				str3 = append(str3, arr2[counter2])
				counter2++
			}
		}
	}
	return strings.Join(str3, "")
}
