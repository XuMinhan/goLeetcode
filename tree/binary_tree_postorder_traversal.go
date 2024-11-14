package main

func postorderTraversal(root *TreeNode) []int {
	var res []int
	var helper func(node *TreeNode)
	helper = func(node *TreeNode) {
		if node == nil {
			return
		}
		helper(node.Left)
		helper(node.Right)
		res = append(res, node.Val)
	}
	helper(root)
	return res
}
