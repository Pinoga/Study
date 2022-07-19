package algo

import . "turing/ds"

func ZipperList(first, second *LinkedListNode) *LinkedListNode {

	if first == nil {
		return second
	}

	if second == nil {
		return first
	}

	var curFirst *LinkedListNode = first
	var curSec *LinkedListNode = second

	for curSec != nil {
		if curFirst.Next == nil {
			curFirst.Next = curSec
			break
		}

		nextFirst := curFirst.Next
		curFirst.Next = curSec

		nextSec := curSec.Next
		curSec.Next = nextFirst

		curSec = nextSec
		curFirst = nextFirst
	}

	return first
}
