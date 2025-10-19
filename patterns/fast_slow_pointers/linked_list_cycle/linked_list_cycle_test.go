package main

import "testing"

func TestHasCycle(t *testing.T) {
	tests := []struct {
		name string
		head *ListNode
		want bool
	}{
		{"case 1: has cycle", buildListWithCycle([]int{3, 2, 0, -4}, 1), true},
		{"case 2: has cycle", buildListWithCycle([]int{1, 2}, 0), true},
		{"case 3: no cycle", buildListWithCycle([]int{1}, -1), false},
		{"case 4: no cycle", buildListWithCycle([]int{1, 2}, -1), false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := hasCycle(test.head)
			if result != test.want {
				t.Errorf("Got: %t, want: %t", result, test.want)
			}
		})
	}
}

func buildListWithCycle(values []int, pos int) *ListNode {
	if len(values) == 0 {
		return nil
	}

	head := &ListNode{
		values[0],
		nil,
	}

	current := head
	var nodes []*ListNode

	nodes = append(nodes, head)

	for i := 1; i < len(values); i++ {
		current.Next = &ListNode{
			values[i],
			nil,
		}
		current = current.Next
		nodes = append(nodes, current)
	}

	if pos >= 0 && pos < len(nodes) {
		current.Next = nodes[pos]
	}

	return head
}
