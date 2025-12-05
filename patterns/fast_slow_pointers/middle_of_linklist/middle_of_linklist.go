// Package middleoflinklist: find the middle of a link list
package middleoflinklist

// https://leetcode.com/problems/middle-of-the-linked-list/description/

type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}
