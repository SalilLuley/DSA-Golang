package leetcode_75

import "fmt"

func StringCompression() {
	s := []byte{'a', 'a', 'b', 'b', 'c', 'c', 'c'}
	fmt.Printf("compress(s): %v\n", compress(s))
}

// Brute force
// func compress(chars []byte) int {
// 	if len(chars) == 0 {
// 		return 0
// 	}
// 	if len(chars) == 1 {
// 		return 1
// 	}
// 	s := ""
// 	currentChar := chars[0]
// 	counter := 1
// 	for i := 1; i < len(chars); i++ {
// 		if i == len(chars)-1 {
// 			counter++
// 			s += string(currentChar) + fmt.Sprint(counter)
// 		} else if chars[i] == currentChar {
// 			counter++
// 			continue
// 		} else {
// 			s += string(currentChar) + fmt.Sprint(counter)
// 			counter = 1
// 			currentChar = chars[i]
// 		}
// 	}
// 	return len(s)
// }

// Time Complexity: O(n)
// Space Complexity: O(1)
func compress(chars []byte) int {
	nu_chars := len(chars)

	if nu_chars == 0 {
		return 0
	}
	if nu_chars == 1 {
		return 1
	}

	writeIndex := 0
	count := 1

	for i := 0; i < len(chars); i++ {
		if i < len(chars)-1 && chars[i] == chars[i+1] {
			count++
		} else {
			chars[writeIndex] = chars[i]
			writeIndex++
			if count > 1 {
				for _, v := range fmt.Sprint(count) {
					chars[writeIndex] = byte(v)
					writeIndex++
				}
			}
			count = 1
		}
	}
	return writeIndex
}
