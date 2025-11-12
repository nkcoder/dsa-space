package middleoflinklist

import (
	"testing"
)

func buildLinkList(vals []int) *ListNode {
	if len(vals) == 0 {
		return nil
	}

	head := &ListNode{
		Val:  vals[0],
		Next: nil,
	}

	current := head

	for i := 1; i < len(vals); i++ {
		node := &ListNode{
			vals[i],
			nil,
		}
		current.Next = node
		current = node
	}

	return head
}

func TestMiddleOfLinkList(t *testing.T) {
	tests := []struct {
		name string
		head *ListNode
		want *ListNode
	}{
		{"case 1: odd nodes", buildLinkList([]int{1, 2, 3, 4, 5}), &ListNode{3, nil}},
		{"case 2: even nodes", buildLinkList([]int{1, 2, 3, 4, 5, 6}), &ListNode{4, nil}},
		{"case 3: only one node", buildLinkList([]int{1}), &ListNode{1, nil}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := middleNode(test.head)
			if result.Val != test.want.Val {
				t.Errorf("Got: %v, want: %v", result, test.want)
			}
		})
	}

}
