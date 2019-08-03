package main

import (
	"strconv"
	"strings"
)

func recoverFromPreorder(S string) *TreeNode {

	levels := make([]int, 1)

	for i := 1; i < len(S); {

		if S[i] == '-' {
			j := i
			for S[j] == '-' {
				j++
			}
			levels = append(levels, j-i)
			i = j
		} else {
			i++
		}
	}

	eles := strings.FieldsFunc(S, func(r rune) bool {
		return r == '-'
	})

	var v func([]string, []int, int) *TreeNode

	v = func(eles []string, levels []int, depth int) *TreeNode {

		if len(eles) == 0 {
			return nil
		}
		val, _ := strconv.Atoi(eles[0])
		root := &TreeNode{Val: val}

		if len(eles) == 1 {
			return root
		}

		nextDepth := depth + 1

		leftStart := 0
		rightStart := len(eles)

		for i, level := range levels {
			if level == nextDepth {
				leftStart = i
				break
			}

		}

		for i := len(levels) - 1; i >= 0; i-- {
			if levels[i] == nextDepth {
				rightStart = i
				break
			}
		}

		if leftStart == rightStart {
			rightStart = len(levels)
		}

		root.Left = v(eles[leftStart:rightStart], levels[leftStart:rightStart], nextDepth)

		if rightStart <= len(levels) {
			root.Right = v(eles[rightStart:], levels[rightStart:], nextDepth)
		} else {
			root.Right = nil
		}

		return root
	}
	return v(eles, levels, 0)

}
