package leetcode_75

import (
	"fmt"
	// "strings"
)

func RemovingStarsFromAString() {
	s := "leet**cod*e"
	fmt.Printf("removeStars(s): %v\n", removeStars(s))
}

//	func removeStars(s string) string {
//		stack := Stack{}
//		for _, v := range s {
//			if v == '*' {
//				if !stack.IsEmpty() {
//					stack.Pop()
//				}
//			} else {
//				stack.Push(string(v))
//			}
//		}
//		return strings.Join(stack.items, "")
//	}
//
//	func (s *Stack) IsEmpty() bool {
//		if len(s.items) == 0 {
//			return true
//		}
//		return false
//	}
//
//	func (s *Stack) Push(data string) {
//		s.items = append(s.items, data)
//	}
//
//	func (s *Stack) Pop() {
//		if s.IsEmpty() {
//			return
//		}
//		s.items = s.items[:len(s.items)-1]
//	}
//
//	type Stack struct {
//		items []string
//	}
//
// Time Complexity: O(n)
// Space Complexity: O(n)
func removeStars(s string) string {
	res := []byte{}
	for _, c := range s {
		if c == '*' {
			res = res[0 : len(res)-1]
		} else {
			res = append(res, byte(c))
		}
	}
	return string(res)
}
