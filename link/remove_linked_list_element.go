package main

func main0() {
	head := NewListNode(1) // 创建链表的头节点
	//head := ListNode{Val: 1}
	head.AddNode(2)
	head.AddNode(3)
	head.AddNode(4)
	removeElements(head, 2).PrintList()
}
func removeElements(head *ListNode, val int) *ListNode {
	dummy := &ListNode{} // 创建一个哑节点，指向 head
	dummy.Next = head
	pre := dummy

	for pre.Next != nil {
		if pre.Next.Val == val {
			pre.Next = pre.Next.Next // 删除节点
		} else {
			pre = pre.Next // 移动指针
		}
	}
	return dummy.Next
}
