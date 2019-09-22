package main

import "math"

func removeZeroSumSublists(head *ListNode) *ListNode {

	mark := make(map[int]*ListNode)

	router := make(map[*ListNode]*ListNode)

	p := head

	dummyHead := &ListNode{Val: math.MaxInt32}
	dummyHead.Next = head
	router[head] = dummyHead
	sum := 0
	for p != nil {

		sum += p.Val

		if prevNode, ok := mark[sum]; ok {
			start := router[prevNode]

			start.Next = p.Next
			p = p.Next

			delete(router, prevNode)

			continue
		} else {
			mark[sum] = p
		}
		if p.Next != nil {
			router[p.Next] = p
		}
		p = p.Next

	}
	return dummyHead.Next
}
