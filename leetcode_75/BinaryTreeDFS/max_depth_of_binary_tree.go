package leetcode_75

import (
	"fmt"
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// createNodes constructs a binary tree from a level-order list representation.
func createNodes(values []interface{}) *TreeNode {
	if len(values) == 0 || values[0] == nil {
		return nil
	}

	root := &TreeNode{Val: values[0].(int)}
	queue := []*TreeNode{root}
	index := 1

	for index < len(values) {
		current := queue[0]
		queue = queue[1:]

		// Assign left child
		if index < len(values) && values[index] != nil {
			current.Left = &TreeNode{Val: values[index].(int)}
			queue = append(queue, current.Left)
		}
		index++

		// Assign right child
		if index < len(values) && values[index] != nil {
			current.Right = &TreeNode{Val: values[index].(int)}
			queue = append(queue, current.Right)
		}
		index++
	}

	return root
}

// printNodes prints the nodes of the binary tree in level-order (BFS)
func printNodes(root *TreeNode) {
	if root == nil {
		fmt.Println("Tree is empty")
		return
	}

	queue := []*TreeNode{root}
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		fmt.Print(current.Val, " ")

		if current.Left != nil {
			queue = append(queue, current.Left)
		}
		if current.Right != nil {
			queue = append(queue, current.Right)
		}
	}
	fmt.Println()
}

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
