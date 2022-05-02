package main

/*
 * @lc app=leetcode.cn id=21 lang=golang
 *
 * [21] 合并两个有序链表
 */

// @lc code=start

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {

	if list1 == nil {

		return list2
	}
	if list2 == nil {
		return list1
	}
	head := &ListNode{0, nil}
	pre := head
	c1 := list1
	c2 := list2

	for {
		if c1.Val > c2.Val {
			pre.Next = c2
			c2 = c2.Next
			pre = pre.Next
			// pre.Next = c1
			// c1 = c1.Next
			// pre = pre.Next
		} else if c1.Val <= c2.Val {
			pre.Next = c1
			c1 = c1.Next
			pre = pre.Next
			// pre.Next = c2
			// c2 = c2.Next
			// pre = pre.Next
		}
		if c1 == nil {
			pre.Next = c2
			break
		}
		if c2 == nil {
			pre.Next = c1

			break
		}

	}
	head = head.Next
	return head
}

// @lc code=end
