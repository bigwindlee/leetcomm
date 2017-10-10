package leetcomm

/**
 * Definition for a binary tree node.
 */
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func Height(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + Max(Height(root.Left), Height(root.Right))
}
