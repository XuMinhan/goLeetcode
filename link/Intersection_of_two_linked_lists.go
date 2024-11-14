package main

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	nodes := map[*ListNode]bool{}
	for headA != nil {
		nodes[headA] = true
		headA = headA.Next
	}
	for headB != nil {
		_, exist := nodes[headB]
		if exist {
			return headB
		}
		headB = headB.Next
	}
	return nil
}
