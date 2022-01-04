package linkedlist

// 160
func GetIntersectionNode(headA, headB *ListNode) *ListNode {
	p1, p2 := headA, headB
	for p1 != p2 {
		if p1 != nil {
			p1 = p1.Next
		} else {
			p1 = headB
		}
		if p2 != nil {
			p2 = p2.Next
		} else {
			p2 = headA
		}
	}
	return p1
}

// 141
func HasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}

	fast, slow := head.Next, head
	for fast != slow && fast != nil {
		fast = fast.Next
		slow = slow.Next
		if fast != nil {
			fast = fast.Next
		}
	}
	return fast == slow
}

// 92 left 从1开始
func ReverseBetween(head *ListNode, left int, right int) *ListNode {
	if head == nil || head.Next == nil || left >= right {
		return head
	}
	result := &ListNode{Next: head}
	father := result
	cur := head
	for i := 1; i < left; i++ {
		cur = cur.Next
		father = father.Next
	}

	p := father
	len := right - left
	for i := 0; i < len+1; i++ {
		son := cur.Next

		cur.Next = father
		father = cur
		cur = son
	}

	p.Next.Next = cur
	p.Next = father
	return result.Next
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func SortList(head *ListNode) *ListNode {
	return mergeSort(head)
}

func mergeSort(head *ListNode) *ListNode {
	fast, slow := head, head.Next
	for fast.Next != nil {
		slow = slow.Next
		fast = fast.Next
	}
	mergeSort(head)
}

func merge(a, b *ListNode) *ListNode {
	if a == nil {
		return b
	}
	if b == nil {
		return a
	}

	sentinel := &ListNode{}
	it := sentinel
	for a != nil && b != nil {
		if a.Val > b.Val {
			it.Next = a
			a = a.Next
		} else {
			it.Next = b
			b = b.Next
		}
	}

	if a != nil {
		it.Next = a
	} else if b != nil {
		it.Next = b
	}
	return sentinel.Next
}
