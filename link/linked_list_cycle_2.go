package main

func detectCycle(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	m := map[*ListNode]bool{}
	for head.Next != nil {
		_, exist := m[head]
		if exist {
			return head
		}
		m[head] = true
		head = head.Next
	}
	return nil
}
