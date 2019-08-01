package main

func delNodes(root *TreeNode, to_delete []int) []*TreeNode {

	mark := make(map[int]bool)

	for _, ele := range to_delete {
		mark[ele] = true
	}

	result := make([]*TreeNode, 0)

	var v func(node *TreeNode, add bool) *TreeNode

	v = func(node *TreeNode, add bool) *TreeNode {

		if node == nil {
			return nil
		} else if mark[node.Val] {
			v(node.Left, true)
			v(node.Right, true)
		} else {
			if add {
				result = append(result, node)
			}
			node.Left = v(node.Left, false)
			node.Right = v(node.Right, false)
			return node
		}

		return nil
	}

	v(root, true)
	return result
}
