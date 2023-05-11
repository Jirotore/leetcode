package leetcode

type ListNode struct {
	Val  int
	Next *ListNode
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	head := l1
	buf := 0

	for {
		if l1 == l2 && l1 != nil {
			buf = l1.Val + buf
		} else {
			buf = l1.Val + l2.Val + buf
		}

		if buf-10 < 0 {
			l1.Val, buf = buf, 0
		} else {
			l1.Val, buf = buf-10, 1
		}

		if l1.Next == nil && l2.Next == nil {
			break
		}

		if l1.Next == nil {
			l1.Next = l2.Next
		}

		if l2.Next == nil {
			l2.Next = l1.Next
		}

		l1, l2 = l1.Next, l2.Next
	}

	if buf != 0 {
		l1.Next = &ListNode{Val: buf}
	}

	return head
}
