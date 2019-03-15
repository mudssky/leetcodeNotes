package main

import "fmt"

/*
025 Reverse Nodes in k-Group
Given a linked list, reverse the nodes of a linked list k at a time and return its modified list.

k is a positive integer and is less than or equal to the length of the linked list. If the number of nodes is not a multiple of k then left-out nodes in the end should remain as it is.

Example:

Given this linked list: 1->2->3->4->5

For k = 2, you should return: 2->1->4->3->5

For k = 3, you should return: 3->2->1->4->5

Note:

Only constant extra memory is allowed.
You may not alter the values in the list's nodes, only nodes itself may be changed.

给出一个链表，每 k 个节点一组进行翻转，并返回翻转后的链表。

k 是一个正整数，它的值小于或等于链表的长度。如果节点总数不是 k 的整数倍，那么将最后剩余节点保持原有顺序。

示例 :

给定这个链表：1->2->3->4->5

当 k = 2 时，应当返回: 2->1->4->3->5

当 k = 3 时，应当返回: 3->2->1->4->5

说明 :

你的算法只能使用常数的额外空间。
你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。
*/
// 这题的难点在于只能使用常数的额外空间。
// 如果使用更多的额外空间的话，用024的申请一个额外数组的方法就可以
// 我们先用这个方法试一下,看看能不能通过
// 解法一,把链表存在数组里,在数组里对节点进行翻转,然后拼接起来
// ListNode 单链表节点
type ListNode struct {
	Val  int
	Next *ListNode
}

func (ln *ListNode) Add(val int) {
	for ln.Next != nil {
		ln = ln.Next
	}
	ln.Next = &ListNode{val, nil}

}
func (ln *ListNode) traversal() {
	tmp := ln
	fmt.Print(tmp.Val)
	for tmp.Next != nil {
		tmp = tmp.Next
		fmt.Print(",", tmp.Val)
	}
	fmt.Println()
}

func traversal(ln *ListNode) {
	for {
		if ln.Next != nil {
			fmt.Print(ln.Val, ",")
		} else {
			fmt.Print(ln.Val, "\n")
			break
		}
		ln = ln.Next
	}
}

// 解法一：时间复杂度O(nlogn),空间复杂度O(n)
// 虽然不满足要求，但是效率却很快
/*
Runtime: 4 ms, faster than 100.00% of Go online submissions for Reverse Nodes in k-Group.
Memory Usage: 4.1 MB, less than 16.67% of Go online submissions for Reverse Nodes in k-Group.
*/
/*func reverseKGroup(head *ListNode, k int) *ListNode {
	nodeList := []*ListNode{}
	tmp := head
	for tmp != nil {
		nodeList = append(nodeList, tmp)
		tmp = tmp.Next
	}
	nodeListLen := len(nodeList)
	// nodeListLen-i >= k 当小于k的时候，剩下不足k的节点就是题中所说的剩余节点
	for i := 0; i <= nodeListLen-k; i += k {
		// 颠倒i到i+k之前的节点
		for l1, l2 := i, i+k-1; l1 < l2; l1, l2 = l1+1, l2-1 {
			nodeList[l1], nodeList[l2] = nodeList[l2], nodeList[l1]
		}
	}
	nodeList = append(nodeList, nil)
	for i := 0; i < nodeListLen; i++ {
		nodeList[i].Next = nodeList[i+1]
	}
	return nodeList[0]
}*/
// 解法二：用栈，可以用大小为k的栈，然而这也不是常数级啊，不断压栈，栈里的节点数量达到k，就全部出栈拼到头节点上
// 如果到结尾了，检查有没有到k个节点，没有的话，按原来的顺序拼接上去，而不是出栈

// 解法三，整体的逻辑是，走k个节点，然后倒着拼上去，这样实现逻辑就可以
func reverseKGroup(head *ListNode, k int) *ListNode {
	if k <= 1 {
		return head
	}
	nodesLen := 0
	for p := head; p != nil; p = p.Next {
		nodesLen++
	}
	rhead := new(ListNode)
	rhead.Next = head
	pre := rhead
	tail, next := head, head
	for ; nodesLen >= k; nodesLen -= k {
		for i := 0; i < k; i++ {
			temp := next.Next
			next.Next = pre.Next
			pre.Next = next
			next = temp
			tail.Next = next
		}
		pre = tail
		tail = next
	}

	return rhead.Next
}

func main() {
	list1 := &ListNode{1, nil}
	list1.Add(2)
	list1.Add(3)
	list1.Add(4)
	list1.Add(5)
	list1.traversal()
	list2 := reverseKGroup(list1, 2)
	list2.traversal()
}
