package main

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	ret := head
	m := map[int]*ListNode{}
	index := 0
	for head != nil {
		m[index] = head
		index++
		head = head.Next
	}
	index--
	index = index - n + 1
	if index == 0 {
		return m[1]
	}
	m[index-1].Next = m[index+1]
	return ret
}
