package Leetcode_75

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func OddEvenList() {
	head := createLinkedList([]int{1, 2, 3, 4, 5})
	oddEvenList(head)
}

func printLinkedList(head *ListNode) {
	current := head
	for current != nil {
		fmt.Printf("current.Val: %v\n", current.Val)
		current = current.Next
	}
	fmt.Println("nil")
}

func createLinkedList(values []int) *ListNode {
	if len(values) == 0 {
		return nil
	}
	head := &ListNode{Val: values[0]}
	current := head

	for _, val := range values[1:] {
		current.Next = &ListNode{Val: val}
		current = current.Next
	}
	return head
}

// Time Complexity - O(n)
// Space Complexity - O(1)
func oddEvenList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	odd := head
	even, evenHead := head.Next, head.Next

	for even != nil && even.Next != nil {
		odd.Next = even.Next
		odd = odd.Next
		even.Next = odd.Next
		even = even.Next
	}

	odd.Next = evenHead
	return head
}
