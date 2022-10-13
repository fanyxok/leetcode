package main

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1.Val == 0 && l1.Next == nil {
		return l2
	} else if l2.Val == 0 && l2.Next == nil {
		return l1
	}
	carry := 0
	head := &ListNode{}
	head_ := head
	for l1 != nil || l2 != nil {
		x := 0
		if l1 != nil {
			x = l1.Val
		}
		y := 0
		if l2 != nil {
			y = l2.Val
		}
		z := x + y + carry
		if z >= 10 {
			z = z - 10
			carry = 1
		} else {
			carry = 0
		}
		head.Next = &ListNode{Val: z}
		head = head.Next
		if l1 != nil {
			l1 = l1.Next
		}
		if l2 != nil {
			l2 = l2.Next
		}
	}
	if carry == 1 {
		head.Next = &ListNode{Val: 1}
	}
	return head_.Next
}
