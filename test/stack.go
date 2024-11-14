package main

import (
	"container/list"
	"fmt"
)

func main1() {
	stack := list.New() // 创建一个新的双向链表作为栈

	// Push 操作
	stack.PushBack(10)
	stack.PushBack(20)
	stack.PushBack(30)
	fmt.Println("Stack size after pushes:", stack.Len())

	// Pop 操作
	top := stack.Back() // 获取最后一个元素
	if top != nil {
		fmt.Println("Popped value:", top.Value)
		stack.Remove(top) // 删除最后一个元素
	}
	fmt.Println("Stack size after pop:", stack.Len())
}
