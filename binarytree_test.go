package leetcomm

import (
	"fmt"
	"math"
	"testing"
)

func TestBinaryTree(t *testing.T) {
	mytree := NewTree([]int{3, 9, 20, math.MinInt32, math.MinInt32, 15, 7})

	t.Errorf("Preorder: %v", PreOrder2Strings(mytree))
}

func PreOrder2Strings(root *TreeNode) []string {
	ret := make([]string, 0)
	if root == nil {
		return ret
	}
	if root.Left == nil && root.Right == nil {
		if root.Val == math.MinInt32 {
			ret = append(ret, "nil")
		} else {
			ret = append(ret, fmt.Sprintf("%d", root.Val))
		}
		return ret
	}

	// Preorder traversal
	ret = append(ret, fmt.Sprintf("%d", root.Val))
	ret = append(ret, PreOrder2Strings(root.Left)...)
	ret = append(ret, PreOrder2Strings(root.Right)...)
	return ret
}

// func LevelOrder2Strings(root *TreeNode) []string {
//   ret := make([]string, 0)
//
// }
