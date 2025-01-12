package leetcode_75

import "fmt"

// Time complexity: O(n)
// Space complexity: O(1)
func Fibonacci() {
	x := 8
	fmt.Printf("fibonacci(): %v\n", fibonacci(x))
}

func fibonacci(x int) int {
	if x == 0 {
		return 0
	} else if x == 1 {
		return 1
	} else if x == 2 {
		return 1
	}
	a, b := 0, 1
	for i := 2; i < x; i++ {
		a, b = b, a+b
	}
	return b
}
