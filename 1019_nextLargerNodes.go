package main

import "container/list"

type ValuePair struct {
	Val   int
	Index int
}

func nextLargerNodes(head *ListNode) []int {

	values := make([]int, 0)

	for head != nil {
		values = append(values, head.Val)
		head = head.Next
	}

	result := make([]int, len(values))
	if len(values) == 0 {
		return result
	}

	stack := list.New()

	stack.PushBack(ValuePair{Val: values[0], Index: 0})

	for i := 1; i < len(values); i++ {

		top := stack.Back()

		if top.Value.(ValuePair).Val < values[i] {

			for top.Value.(ValuePair).Val < values[i] {

				result[top.Value.(ValuePair).Index] = values[i]

				stack.Remove(top)
				if stack.Len() == 0 {
					break
				}

				top = stack.Back()
			}
		}
		stack.PushBack(ValuePair{Val: values[i], Index: i})
	}
	return result
}
