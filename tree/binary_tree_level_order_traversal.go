package main

import (
	"container/list"
	"fmt"
)

func main0() {
	node := &TreeNode{
		Val:   0,
		Left:  nil,
		Right: nil,
	}
	order := levelOrder(node)
	fmt.Println(order)
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	l := list.New()
	l.PushBack(root)
	ret := [][]int{}
	for l.Len() != 0 {
		newL := list.New()
		thisLevel := []int{}
		for l.Len() != 0 {
			front := l.Front()
			node := front.Value.(*TreeNode)
			thisLevel = append(thisLevel, node.Val)
			l.Remove(front)
			if node.Left != nil {
				newL.PushBack(front.Value.(*TreeNode).Left)
			}
			if node.Right != nil {
				newL.PushBack(front.Value.(*TreeNode).Right)
			}
		}
		ret = append(ret, thisLevel)
		l.PushFrontList(newL)
	}
	return ret
}
