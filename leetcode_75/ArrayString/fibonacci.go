package leetcode_75

import "fmt"

// https://leetcode.com/problems/fibonacci-number/description/
// Time complexity: O(n)
// Space complexity: O(1)
func Fibonacci() {
	x := 8
	fmt.Printf("fibonacci(): %v\n", fib(x))
}

func fib(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	}
	a, b := 0, 1
	for i := 1; i < n; i++ {
		a, b = b, a+b
	}
	return b
}
