package Leetcode_75

func ReverseLinkedList() {
	head := createLinkedList([]int{1, 2, 3, 4, 5})
	printLinkedList(reverseList(head))
}

// Question - 1->2->3->4->5
// None ->  1 -> 2 -> 3 -> 4 -> 5
// p		h    h.n
//			p	 h	  h.n

// Time Complexity O(n)
// Space Complexity O(1)
func reverseList(head *ListNode) *ListNode {
	curr := head
	var prev *ListNode
	for curr != nil {
		nextNode := curr.Next
		curr.Next = prev
		prev = curr
		curr = nextNode
	}
	return prev
}
