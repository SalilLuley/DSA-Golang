package Leetcode_75

import (
	"math"
)

func MaximumTwinSumOfLinkedList() {
	head := createLinkedList([]int{1, 100000})
	pairSum(head)
}

func pairSum(head *ListNode) int {
	max := 0
	slow, fast := head, head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	var prev *ListNode

	for slow != nil {
		nextNode := slow.Next
		slow.Next = prev
		prev = slow
		slow = nextNode
	}

	for head != nil && prev != nil {
		max = int(math.Max(float64(max), float64(head.Val)+float64(prev.Val)))
		head = head.Next
		prev = prev.Next
	}
	return max
}
