package leetcode_75

import (
	"fmt"
	"math"
)

// Time complexity - O(n)
func MaxDepth() {
	values := []interface{}{3, 9, 20, nil, nil, 15, 7}
	nodes := createNodes(values)
	result := maxDepth(nodes)
	fmt.Printf("result: %v\n", result)

}

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + int(math.Max(float64(maxDepth(root.Left)), float64(maxDepth(root.Right))))
}
