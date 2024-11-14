package main

import "fmt"

func isSymmetric(root *TreeNode) bool {
	var res1 []int
	var helper1 func(node *TreeNode)
	helper1 = func(node *TreeNode) {
		if node == nil {
			res1 = append(res1, 101)
			return
		}
		res1 = append(res1, node.Val)
		helper1(node.Left)
		helper1(node.Right)
	}
	helper1(root)

	var res2 []int
	var helper2 func(node *TreeNode)
	helper2 = func(node *TreeNode) {
		if node == nil {
			res2 = append(res2, 101)
			return
		}
		res2 = append(res2, node.Val)
		helper2(node.Right)
		helper2(node.Left)
	}
	helper2(root)
	fmt.Println(res1)
	fmt.Println(res2)
	for i := range res1 {
		if res1[i] != res2[i] {
			return false
		}
	}
	return true
}
