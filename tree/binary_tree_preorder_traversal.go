package main

func preorderTraversal(root *TreeNode) []int {
	var res []int
	var helper func(node *TreeNode)
	helper = func(node *TreeNode) {
		if node == nil {
			return
		}
		res = append(res, node.Val)
		helper(node.Left)
		helper(node.Right)
	}
	helper(root)
	return res
}
