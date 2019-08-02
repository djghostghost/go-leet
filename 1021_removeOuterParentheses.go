package main

import "strings"

func removeOuterParentheses(S string) string {

	mark := make([]int, len(S)+1)

	zeroIndexs := make([]int, 1)

	for i, ch := range S {

		if ch == '(' {
			mark[i+1] = mark[i] + 1
		} else {
			mark[i+1] = mark[i] - 1
		}

		if mark[i+1] == 0 {
			zeroIndexs = append(zeroIndexs, i+1)
		}
	}

	var sb strings.Builder

	for i := 0; i < len(zeroIndexs)-1; i++ {
		sb.WriteString(S[zeroIndexs[i]+1 : zeroIndexs[i+1]-1])
	}

	return sb.String()
}
