package main

import (
	"container/list"
	"strings"
)

func reverseParentheses(s string) string {

	stack := list.New()

	var sb strings.Builder
	for _, ch := range s {

		if ch == ')' {
			sb.Reset()
			for stack.Len() != 0 {
				top := stack.Back()
				//往回写
				if top.Value.(int32) == '(' {
					str := sb.String()
					for j := 0; j < len(str); j++ {
						stack.PushBack(int32(str[j]))
					}
					stack.Remove(top)
					break
				} else {
					sb.WriteRune(rune(top.Value.(int32)))
					stack.Remove(top)

				}
			}
		} else {
			stack.PushBack(ch)
		}
	}
	sb.Reset()
	for top := stack.Front(); top != nil; top = top.Next() {

		sb.WriteRune(rune(top.Value.(int32)))
	}
	return sb.String()
}
