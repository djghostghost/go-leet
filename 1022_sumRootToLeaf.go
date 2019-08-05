package main

func sumRootToLeaf(root *TreeNode) int {

	result := 0

	var v func(*TreeNode, int)

	v = func(node *TreeNode, subResult int) {

		if node == nil {
			return
		}
		subResult = subResult*2 + node.Val
		if node.Left == nil && node.Right == nil {
			result += subResult
			return
		}

		v(node.Left, subResult)
		v(node.Right, subResult)

	}

	v(root, 0)
	return result
}
