package algo

import (
	. "turing/ds"
)

func Reverse(head *LinkedListNode) *LinkedListNode {

	if head == nil || head.Next == nil {
		return head
	}

	var previous *LinkedListNode
	var Next *LinkedListNode
	current := head

	for current != nil {

		Next = current.Next
		current.Next = previous

		previous = current
		current = Next
	}

	return previous
}
