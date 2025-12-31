package array

import "fmt"

func StringCompress() {
	chars := []byte{'a', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b'}
	fmt.Printf("compress %v \n", compress(chars))
}

// 0 1 2
// a a b b c c c
//          s
//          e

func compress(chars []byte) int {
	return 0
}
