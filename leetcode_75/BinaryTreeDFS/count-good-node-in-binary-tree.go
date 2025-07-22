package leetcode_75

import (
	"fmt"
	"math"
)

func GoodNode() {
	root1 := []interface{}{3, 1, 4, 3, nil, 1, 5}
	head1 := createNodes(root1)
	max := head1.Val
	fmt.Printf("goodNodes(root): %v\n", goodNodes(head1, max))
}

func goodNodes(root *TreeNode, max int) int {
	return dfs(root, max)
}

func dfs(root *TreeNode, max int) int {
	if root == nil {
		return 0
	}

	good := 0
	if root.Val >= max {
		good = 1
	}

	newMax := int(math.Max(float64(root.Val), float64(max)))
	good += dfs(root.Left, newMax)
	good += dfs(root.Right, newMax)

	return good
}
