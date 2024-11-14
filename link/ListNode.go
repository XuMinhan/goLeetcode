package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func NewListNode(Val int) *ListNode {
	return &ListNode{Val: Val, Next: nil}
}

func (node *ListNode) AddNode(Val int) {
	for node.Next != nil {
		node = node.Next
	}
	node.Next = NewListNode(Val)
}

func (node *ListNode) PrintList() {
	for node != nil {
		fmt.Printf("%d -> ", node.Val)
		node = node.Next
	}
	fmt.Println(nil)
}

func main_() {
	head := NewListNode(1) // 创建链表的头节点
	//head := ListNode{Val: 1}
	head.AddNode(2)
	head.AddNode(3)
	head.AddNode(4)

	head.PrintList() // 输出：1 -> 2 -> 3 -> 4 -> nil
}
