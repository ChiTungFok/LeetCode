package goodnodes

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func goodNodes(root *TreeNode) int {
	return findGoodNodes(root, root.Val)
}

func findGoodNodes(root *TreeNode, maxVal int) int {
	var v, l, r int
	if root.Val >= maxVal {
		v++
	}
	if root.Left != nil {
		l = findGoodNodes(root.Left, max(maxVal, root.Val))
	}
	if root.Right != nil {
		r = findGoodNodes(root.Right, max(maxVal, root.Val))
	}
	return v + l + r
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
