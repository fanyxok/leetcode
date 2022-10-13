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

// nums 是head的子集， 遍历head检查nums即可
// 遇到一个component计数，还是离开一个component计数，
// 这里遍历结束时可能还没有离开，所以需要额外判断一次计数
func numComponents(head *ListNode, nums []int) int {
	set := make(map[int]struct{}, len(nums))
	for _, v := range nums {
		set[v] = struct{}{}
	}
	n := 0
	nodeInSet := false
	for ; head != nil; head = head.Next {
		if _, ok := set[head.Val]; ok {
			nodeInSet = true
		} else if nodeInSet {
			n++
			nodeInSet = false
		}
	}
	if nodeInSet {
		n++
	}
	return n
}
