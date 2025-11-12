// Package linkedlistcycle check if there is cycle in a linked list
package linkedlistcycle

// https://leetcode.com/problems/linked-list-cycle/
// Fast slow pointers: need to pay attention to the for condition (nil pointer reference)

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next

		if slow == fast {
			return true
		}
	}
	return false
}
