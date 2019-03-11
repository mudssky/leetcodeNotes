package main

/*
Given a linked list, swap every two adjacent nodes and return its head.

You may not modify the values in the list's nodes, only nodes itself may be changed.

Example:

Given 1->2->3->4, you should return the list as 2->1->4->3.

给定一个链表，两两交换其中相邻的节点，并返回交换后的链表。

你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。

示例:

给定 1->2->3->4, 你应该返回 2->1->4->3.

*/
import "fmt"

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

// 解法一：用另一个切片帮助进行存储，遍历数组进行翻转交换位置，再遍历一遍，把节点的next连接起来
// 时间复杂度O(n) 空间复杂度O(n)
/*func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	nodeList := []*ListNode{}
	for head != nil {
		nodeList = append(nodeList, head)
		head = head.Next
	}
	for i := 0; i < len(nodeList)-1; i += 2 {
		nodeList[i], nodeList[i+1] = nodeList[i+1], nodeList[i]
	}
	nodeList = append(nodeList, nil)
	for i := 0; i < len(nodeList)-1; i++ {
		nodeList[i].Next = nodeList[i+1]
	}
	return nodeList[0]
}*/

// 解法二：递归求解
func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	tmp := head.Next
	head.Next = swapPairs(tmp.Next)
	tmp.Next = head
	return tmp
}
func main() {
	list1 := ListNode{1, nil}
	list1.Add(2)
	list1.Add(3)
	list1.Add(4)
	list1.traversal()
	list2 := swapPairs(&list1)
	list2.traversal()
}
