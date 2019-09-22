package main

import (
	"container/list"
	"math"
)

type NodeInfo struct {
	Level int
	Node  *TreeNode
}

func maxLevelSum(root *TreeNode) int {

	queue := list.New()

	queue.PushBack(NodeInfo{Level: 1, Node: root})

	mark := make(map[int]int)

	for queue.Len() != 0 {
		top := queue.Front()

		level := top.Value.(NodeInfo)
		mark[level.Level] += level.Node.Val

		if level.Node.Left != nil {
			queue.PushBack(NodeInfo{Level: level.Level + 1, Node: level.Node.Left})
		}

		if level.Node.Right != nil {
			queue.PushBack(NodeInfo{Level: level.Level + 1, Node: level.Node.Right})
		}

		queue.Remove(top)
	}

	result := 0

	maxSum := math.MinInt32
	for key, val := range mark {
		if maxSum < val {
			maxSum = val
			result = key
		}
	}
	return result
}
