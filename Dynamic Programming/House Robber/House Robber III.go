package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rob(root *TreeNode) int {
	var _rob func(root *TreeNode) (rob, notRob int)
	_rob = func(root *TreeNode) (rob, notRob int) {
		if root == nil {
			return 0, 0
		}
		/*
				   4
				  1
				 2
				3
				  3 0
			 	  2 3
				  4 2
		*/
		robLeftMax, notRobLeftMax := _rob(root.Left)
		robRightMax, notRobRightMax := _rob(root.Right)
		rob = notRobLeftMax + notRobRightMax + root.Val
		notRob = max(robLeftMax, notRobLeftMax) + max(robRightMax, notRobRightMax)
		return rob, notRob
	}
	r, notRob := _rob(root)
	return max(r, notRob)
}
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
