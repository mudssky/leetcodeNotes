package main

import "fmt"

/*
Problem:
Given a linked list, remove the n-th node from the end of list and return its head.

Example:

Given linked list: 1->2->3->4->5, and n = 2.

After removing the second node from the end, the linked list becomes 1->2->3->5.
Note:

Given n will always be valid.

Follow up:

Could you do this in one pass?

给定一个链表，删除链表的倒数第 n 个节点，并且返回链表的头结点。

示例：

给定一个链表: 1->2->3->4->5, 和 n = 2.

当删除了倒数第二个节点后，链表变为 1->2->3->5.
说明：

给定的 n 保证是有效的。

进阶：

你能尝试使用一趟扫描实现吗？
*/
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

func (ln *ListNode) traversalFrom(begin *ListNode) {
	fmt.Print(begin.Val)
	for begin.Next != nil {
		begin = begin.Next
		fmt.Print(",", begin.Val)
	}
	fmt.Println()

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

//Add 单向列表添加元素在最后面
func (ln *ListNode) Add(val int) {
	for ln.Next != nil {
		ln = ln.Next
	}
	ln.Next = &ListNode{val, nil}

}

// 解法一，用另一个指针切片保存每个节点的指针，然后从这个切片中定位进行操作
// 时间复杂度 O(n),空间复杂度O(1)
/*
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	nodeList := []*ListNode{}
	nodeList = append(nodeList, head)
	for head.Next != nil {
		head = head.Next
		nodeList = append(nodeList, head)
	}
	nodeListLen := len(nodeList)
	if n == 1 {
		if nodeListLen == 1 {
			return nil
		} else {
			// 修改倒数第二个的next为nil即可
			nodeList[nodeListLen-n-1].Next = nil
			return nodeList[0]
		}
	} else if n == nodeListLen {
		nodeList[0].Next = nil
		return nodeList[1]
	} else {
		nodeList[nodeListLen-n-1].Next = nodeList[nodeListLen-n].Next
		return nodeList[0]
	}
}*/
// 解法二：用于定位的话实际上一个指针的切片有点多余了，浪费很多空间
// 只需要两个指针即可，第一个指针向前先移动n+1步，第二个指针再从开头开始出发
// 这样当第一个指针遍历到结尾，到nil节点的时候，第二个指针正是倒数第n个节点前一个节点。

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	first := head
	second := head
	for i := 0; i < n+1; i++ {
		// 有个问题，第一个节点可能走过头，比如n是1，链表长度是1的情况，因为起始点已经+1还需要走两次
		// 会碰到nil而走不下去。
		// 所以说要判断一下，当first走到末尾的时候，如果此时i==n-1就是那种情况,说明要删除的是头节点
		first = first.Next
		if first == nil && i == n-1 {
			res := head.Next
			head.Next = nil
			return res
		}
	}
	// 此外如果n==1，删除尾节点的时候也要特殊处理，所以说不如分开判断
	for {
		// 当第一个节点走到末尾，说明第二个节点现在在该删除的节点之前的一个节点
		if first == nil {
			if n == 1 {
				second.Next = nil
				return head
				// 除此之外就是在中间的情况，让前一个节点的指向下一个节点即可
			} else {
				second.Next = second.Next.Next
				return head
			}

		}
		first = first.Next
		second = second.Next

	}

}
func main() {
	input := &ListNode{1, nil}
	input.Add(2)
	input.traversal()
	output := removeNthFromEnd(input, 1)
	output.traversal()
}
