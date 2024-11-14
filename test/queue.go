package main

import (
	"container/list"
	"fmt"
)

func main0() {
	queue := list.New() // 创建一个新的双向链表作为队列

	// Enqueue 操作
	queue.PushBack(10)
	queue.PushBack(20)
	queue.PushBack(30)
	fmt.Println("Queue size after enqueues:", queue.Len())

	// Dequeue 操作
	front := queue.Front() // 获取第一个元素
	if front != nil {
		fmt.Println("Dequeued value:", front.Value)
		queue.Remove(front) // 删除第一个元素
	}
	fmt.Println("Queue size after dequeue:", queue.Len())
}
