package leetcode_75

import (
	"fmt"
	"slices"
)

func LeafSimilar() {
	root1 := []interface{}{3, 5, 1, 6, 2, 9, 8, nil, nil, 7, 4}
	root2 := []interface{}{3, 5, 1, 6, 7, 4, 2, nil, nil, nil, nil, nil, nil, 9, 8}

	head1 := createNodes(root1)
	head2 := createNodes(root2)
	fmt.Printf("leafSimilar(head1, head2): %v\n", leafSimilar(head1, head2))
}

//Time Complexity  O(N + M)
//Space Complexity O(L1 + L2)

func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	var leaf1 []int
	var leaf2 []int

	getLeaf(root1, &leaf1)
	getLeaf(root2, &leaf2)
	return leaf1 != nil && leaf2 != nil && slices.Equal(leaf1, leaf2)
}

func getLeaf(root *TreeNode, result *[]int) {
	if root.Left == nil && root.Right == nil {
		*result = append(*result, root.Val)
		return
	}

	if root.Left != nil {
		getLeaf(root.Left, result)
	}

	if root.Right != nil {
		getLeaf(root.Right, result)
	}
}
