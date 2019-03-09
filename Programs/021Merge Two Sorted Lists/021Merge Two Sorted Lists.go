package main

/*
Merge two sorted linked lists and return it as a new list. The new list should be made by splicing together the nodes of the first two lists.

Example:

Input: 1->2->4, 1->3->4
Output: 1->1->2->3->4->4
将两个有序链表合并为一个新的有序链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。

示例：

输入：1->2->4, 1->3->4
输出：1->1->2->3->4->4
*/

/*
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

// 解法一：同时遍历遍历两个链表，比较两个链表节点大小，分别拼到新链表后面
// 时间复杂度，遍历玩两个链表，所以是O(m+n) 空间复杂度，O(m+n)
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	// 如果其中一个是空链表，那么新链表就是另一个链表
	if l1 == nil {
		return l2
	} else if l2 == nil {
		return l1
	}
	// 设置一个起始点值为0（其实值无所谓），然后遍历的时候，慢慢在后面把节点拼接上去
	zeroNode := &ListNode{0, nil}
	// 临时指针，用于帮助构造新链表
	// 比较当前节点的Val大小，小的拼接到新链表上，并且指针继续往后移
	tmpNode := zeroNode
	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			tmpNode.Next = &ListNode{l1.Val, nil}
			tmpNode = tmpNode.Next
			l1 = l1.Next
		} else {
			tmpNode.Next = &ListNode{l2.Val, nil}
			tmpNode = tmpNode.Next
			l2 = l2.Next
		}
	}
	// 结束上面的流程以后，会有一条链表已经全部插入完毕
	// 判断，并且把另一条链表剩下的全部插入
	if l1 == nil {
		for l2 != nil {
			tmpNode.Next = &ListNode{l2.Val, nil}
			tmpNode = tmpNode.Next
			l2 = l2.Next
		}
	} else if l2 == nil {
		for l1 != nil {
			tmpNode.Next = &ListNode{l1.Val, nil}
			tmpNode = tmpNode.Next
			l1 = l1.Next
		}
	}
	// 返回头节点
	return zeroNode.Next
}
func main() {

}
