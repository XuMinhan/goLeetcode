package main

func main2() {
	head := NewListNode(1) // 创建链表的头节点
	head.AddNode(2)
	head.AddNode(3)
	head.AddNode(4)
	//head := ListNode{Val: 1}
	head.PrintList() // 输出：1 -> 2 -> 3 -> 4 -> nil
	(swapPairs(head)).PrintList()
}

/*
Input: head = [1,2,3,4]

1 -> 2 -> 3 -> 4 -> 5 -> 6. pre nil ; cur 1   ; next 3
2 -> 1 -> 3 -> 4 -> 5 -> 6. pre 1   ; cur 3   ; next 5
2 -> 1 -> 4 -> 3 -> 5 -> 6. pre 3   ; cur 5   ; next nil
2 -> 1 -> 4 -> 3 -> 6 -> 5. pre 5   ; cur nil ; next nil

Output: [2,1,4,3]
*/

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	ret := head.Next
	var pre *ListNode
	cur := head
	for cur != nil {
		if cur.Next == nil {
			pre.Next = cur
			return ret
		}
		curNext := cur.Next
		next := cur.Next.Next
		if pre != nil {
			pre.Next = curNext
		}
		curNext.Next = cur
		cur.Next = next

		pre = cur
		cur = next
	}
	return ret
}
