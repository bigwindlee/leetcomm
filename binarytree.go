package leetcomm

import (
	"math"
)

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

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	return isSymmetricHelp(root.Left, root.Right)
}

func isSymmetricHelp(left, right *TreeNode) bool {
	if left == nil || right == nil {
		return left == right
	}
	return left.Val == right.Val && isSymmetricHelp(left.Left, right.Right) && isSymmetricHelp(left.Right, right.Left)
}

func Size(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return 1 + Size(root.Left) + Size(root.Right)
}

func IsSubtree(s *TreeNode, t *TreeNode) bool {
	if t == nil {
		return true
	}
	if s == nil {
		return false
	}
	q := make([]*TreeNode, 0)
	q = append(q, s)
	for len(q) > 0 {
		deq := q[0]
		if IsSameTree(deq, t) {
			return true
		}
		q = q[1:]
		if deq.Left != nil {
			q = append(q, deq.Left)
		}
		if deq.Right != nil {
			q = append(q, deq.Right)
		}
	}
	return false
}

func IsSameTree(p, q *TreeNode) bool {
	if p == nil || q == nil {
		return p == q
	}
	return p.Val == q.Val && IsSameTree(p.Left, q.Left) && IsSameTree(p.Right, q.Right)
}

/* Determine if it is AVL tree */
func IsHeightBalanced(root *TreeNode) bool {
	var height int
	return isHeightBalancedHelp(root, &height)
}

/* Remember height avoiding calculating height in each recursive call
   Time Complexity: O(n)*/
func isHeightBalancedHelp(root *TreeNode, height *int) bool {
	/* lh --> Height of left subtree
	   rh --> Height of right subtree */
	lh, rh := 0, 0

	/* l will be true if left subtree is balanced
	   and r will be true if right subtree is balanced */
	l, r := false, false

	if root == nil {
		*height = 0
		return true
	}

	l = isHeightBalancedHelp(root.Left, &lh)
	r = isHeightBalancedHelp(root.Right, &rh)
	*height = Max(lh, rh) + 1
	/* In an AVL tree, the heights of the two child subtrees of any node differ by at most one */
	if Abs(lh, rh) >= 2 {
		return false
	}

	return l && r
}

func NewTree(arr []int) *TreeNode {
	if len(arr) == 0 {
		return nil
	}
	root := &TreeNode{arr[0], nil, nil}
	queue := make([]**TreeNode, 0)
	queue = append(queue, &root.Left)
	queue = append(queue, &root.Right)

	for i := 1; i < len(arr); i++ {
		deq := queue[0]
		queue = queue[1:]
		if arr[i] == math.MinInt32 {
			*deq = nil
		} else {
			*deq = &TreeNode{arr[i], nil, nil}
			queue = append(queue, &((*deq).Left))
			queue = append(queue, &((*deq).Right))
		}
	}

	return root
}
