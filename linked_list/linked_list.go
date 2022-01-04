package linkedlist

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func New(a []int) *ListNode {
	if len(a) == 0 {
		return nil
	}

	head := &ListNode{}
	cur := head
	for _, v := range a {
		cur.Next = &ListNode{Val: v}
		cur = cur.Next
	}
	return head.Next
}

func NewWith(a []int, b *ListNode) *ListNode {
	if len(a) == 0 {
		return nil
	}

	head := &ListNode{}
	cur := head
	for _, v := range a {
		cur.Next = &ListNode{Val: v}
		cur = cur.Next
	}

	if b != nil {
		cur.Next = b
	}

	return head.Next
}

func NewCycle(a []int) *ListNode {
	if len(a) == 0 {
		return nil
	}

	head := &ListNode{}
	cur := head
	for _, v := range a {
		cur.Next = &ListNode{Val: v}
		cur = cur.Next
	}

	cur.Next = head.Next
	return head.Next
}

func (l *ListNode) Print() (res string) {
	for l != nil {
		res += fmt.Sprintf("%v ", l.Val)
		l = l.Next
	}
	return
}

func (l *ListNode) Eq(o *ListNode) bool {
	for l != nil && o != nil {
		if l.Val != o.Val {
			return false
		}
		l = l.Next
		o = o.Next
	}

	return l == o
}
