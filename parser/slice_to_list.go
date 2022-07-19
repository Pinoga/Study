package parser

import (
	"fmt"
	. "turing/ds"
)

func PrintList(head *LinkedListNode) {
	for cur := head; cur != nil; cur = cur.Next {
		fmt.Printf("%v ", cur.Value)
	}
}

func ParseSlice(s []int) *LinkedListNode {

	if len(s) == 0 {
		return nil
	}

	head := &LinkedListNode{
		Next:  nil,
		Value: NodeValue(s[0]),
	}

	current := head

	for i := 1; i < len(s); i++ {
		new := &LinkedListNode{
			Next:  nil,
			Value: NodeValue(s[i]),
		}
		current.Next = new
		current = new
	}
	return head
}
