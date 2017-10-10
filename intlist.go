package leetcomm

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func (pl *ListNode) String() string {
	var ret []int
	for p := pl; p != nil; p = p.Next {
		ret = append(ret, p.Val)
	}
	return fmt.Sprintf("%v", ret)
}

func GetList(l []int) *ListNode {
	var ret, last *ListNode
	for _, value := range l {
		if ret == nil {
			last = &ListNode{value, nil}
			ret = last
		} else {
			last.Next = &ListNode{value, nil}
			last = last.Next
		}
	}
	return ret
}

func EqualList(a, b *ListNode) bool {
	pa := a
	pb := b
	for pa != nil && pb != nil {
		if pa.Val != pb.Val {
			return false
		}
		pa = pa.Next
		pb = pb.Next
	}
	if pa == nil && pb == nil {
		return true
	} else {
		return false
	}
}
