package main

import "fmt"

func main1() {
	nodeLeft := &TreeNode{
		Val:   1,
		Left:  nil,
		Right: nil,
	}
	nodeRight := &TreeNode{
		Val:   2,
		Left:  nil,
		Right: nil,
	}
	node := &TreeNode{
		Val:   0,
		Left:  nodeLeft,
		Right: nodeRight,
	}

	order := rightSideView(node)
	fmt.Println(order)
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rightSideView(root *TreeNode) []int {
	right := []int{}
	var helper func(node *TreeNode)
	level := 0
	maxLevel := 0
	helper = func(node *TreeNode) {
		if node == nil {
			return
		}
		if level >= maxLevel {
			right = append(right, node.Val)
			maxLevel++
		}
		level++
		helper(node.Right)
		helper(node.Left)
		level--
	}
	helper(root)
	return right
}
