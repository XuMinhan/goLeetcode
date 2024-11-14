package main

func main1() {
	head := NewListNode(1) // 创建链表的头节点
	head.AddNode(2)
	head.AddNode(3)
	head.AddNode(4)

	head.PrintList() // 输出：1 -> 2 -> 3 -> 4 -> nil
	newList := reverseList(head)
	newList.PrintList()

}

func reverseList(head *ListNode) *ListNode {
	var pre *ListNode
	var next *ListNode
	var cur *ListNode
	cur = head
	for cur != nil {
		next = cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}
