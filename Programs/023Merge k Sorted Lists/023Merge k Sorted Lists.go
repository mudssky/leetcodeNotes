package main

import (
	"fmt"
)

/*
Merge k sorted linked lists and return it as one sorted list. Analyze and describe its complexity.

Example:

Input:
[
  1->4->5,
  1->3->4,
  2->6
]
Output: 1->1->2->3->4->4->5->6

合并 k 个排序链表，返回合并后的排序链表。请分析和描述算法的复杂度。

示例:

输入:
[
  1->4->5,
  1->3->4,
  2->6
]
输出: 1->1->2->3->4->4->5->6
*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
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

// 解法一，以前写过一个拼接两个有序链表的函数，只要多次调用这个函数即可，两个两个排

/*func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
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

func mergeKLists(lists []*ListNode) *ListNode {
	listsLen := len(lists)
	if listsLen == 0 {
		return nil
	}
	if listsLen == 1 {
		return lists[0]
	}
	res := mergeTwoLists(lists[0], lists[1])
	if listsLen == 2 {
		return res
	}
	for i := 2; i < listsLen; i++ {
		res = mergeTwoLists(res, lists[i])
	}
	return res
}*/
// 解法二，按照以前的思路，比较最小的部分，就用一个求最小的函数代替。当链表为空的时候直接移出数组
// 时间复杂度 遍历n个长为m链表加上每次遍历最小值遍历n次，O(n*m)
// 空间复杂度，大约为一个新链表的长度 O(n*m)
// getMin 在一个以上ListNode节点中求出最小值所在的节点并返回对应的数组下标
/*func getMin(lists []*ListNode) int {
	listsLen := len(lists)
	if listsLen == 1 {
		return 0
	}
	min := 0
	for i := 1; i < listsLen; i++ {
		if lists[i].Val < lists[min].Val {
			min = i
		}
	}
	return min
}

// 删除指定下标的链表
func deleteEle(lists *[]*ListNode, i int) {
	if len(*lists) == 0 {
		*lists = nil
	}
	if len(*lists) == 2 {
		// 由于指针运算符的优先级比较低，所以先进行取值要用括号括起来操作	优先级：()>[]>*
		if i == 1 {
			*lists = (*lists)[:1]
		} else if i == 0 {
			*lists = (*lists)[1:]

		}
		return
	}
	*lists = append((*lists)[:i], (*lists)[i+1:]...)
}

// 去除所有空链表
func deleteNil(lists *[]*ListNode) (res []*ListNode) {
	for i := 0; i < len(*lists); i++ {
		if (*lists)[i] != nil {
			res = append(res, (*lists)[i])
		}
	}
	return
}

func mergeKLists(lists []*ListNode) *ListNode {
	// 因为测试值中会故意给出元素时空链表的选项，所以首先清除空链表
	lists = deleteNil(&lists)
	// 清除完若果结果是空，因为最后返回root.Next所以包含了这个情况，不需要特别判断传入空数组的情况了
	root := &ListNode{0, nil}
	tmp := root
	// 退出条件是当所有链表变成空链表，因为变成空链表会经过删除逻辑，也就是lists列表为空
	for len(lists) > 0 {
		// 不断获取最小值赋值给新节点
		minPos := getMin(lists)
		tmp.Next = &ListNode{lists[minPos].Val, nil}
		tmp = tmp.Next
		lists[minPos] = lists[minPos].Next

		if lists[minPos] == nil {
			deleteEle(&lists, minPos)
		}
	}
	return root.Next
}*/
// 解法三 优化解法一，通过分治法，把合并k个划分为合并 k/2 不断划分，直到两两合并。相比解法一，大大减少了合并的次数
/*
Runtime: 16 ms, faster than 64.37% of Go online submissions for Merge k Sorted Lists.
Memory Usage: 6.2 MB, less than 13.04% of Go online submissions for Merge k Sorted Lists.
*/
/*func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
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
func mergeKLists(lists []*ListNode) *ListNode {
	listsLen := len(lists)
	// 若传入的列表为空，返回nil
	if listsLen == 0 {
		return nil
	}
	// 退出条件是合并到列表中只有一个元素
	for listsLen > 1 {
		// listsLen+1/2
		// 这是为了当n为奇数的时候，k能始终从后半段开始，比如当n=5时，那么此时k=3，则0和3合并，1和4合并，最中间的2空出来。
		// 当n是偶数的时候，加1也不会有影响，比如当n=4时，此时k=2，那么0和2合并，1和3合并，完美解决问题
		k := (listsLen + 1) / 2
		for i := 0; i < listsLen/2; i++ {
			lists[i] = mergeTwoLists(lists[i], lists[i+k])
		}
		listsLen = k
	}
	return lists[0]
}*/
// 解法四 使用最小堆
// 参考别人的代码如下
/*
Runtime: 8 ms, faster than 100.00% of Go online submissions for Merge k Sorted Lists.
Memory Usage: 5.4 MB, less than 65.22% of Go online submissions for Merge k Sorted Lists.
*/

type NodeHeap []*ListNode

func (h NodeHeap) Len() int           { return len(h) }
func (h NodeHeap) Less(i, j int) bool { return h[i].Val <= h[j].Val }
func (h NodeHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *NodeHeap) Push(x *ListNode) {
	*h = append(*h, x)
	h.up(h.Len() - 1)
}

func (h *NodeHeap) Pop() *ListNode {
	e := h.Len() - 1
	h.Swap(0, e)
	h.down(0, e)

	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func (h *NodeHeap) up(j int) {
	for {
		i := (j - 1) / 2 // parent
		if i == j || !h.Less(j, i) {
			break
		}
		h.Swap(i, j)
		j = i
	}
}

func (h *NodeHeap) down(i0, n int) bool {
	i := i0
	for {
		j1 := 2*i + 1
		if j1 >= n || j1 < 0 { // j1 < 0 after int overflow
			break
		}
		j := j1 // left child
		if j2 := j1 + 1; j2 < n && h.Less(j2, j1) {
			j = j2 // = 2*i + 2  // right child
		}
		if !h.Less(j, i) {
			break
		}
		h.Swap(i, j)
		i = j
	}
	return i > i0
}

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0 {
		return nil
	}
	if len(lists) == 1 {
		return lists[0]
	}

	heap := make(NodeHeap, 0, len(lists))

	for i := 0; i < len(lists); i++ {
		if lists[i] != nil {
			heap.Push(lists[i])
		}
	}

	var head *ListNode
	var tail *ListNode
	var next *ListNode

	for heap.Len() > 0 {
		if tail != nil {
			next = heap.Pop()
			if next.Next != nil {
				heap.Push(next.Next)
			}
			tail.Next = next
			tail = next
		} else {
			head = heap.Pop()
			tail = head
			if head.Next != nil {
				heap.Push(head.Next)
			}
		}
	}

	return head
}

// 做了这题感觉很多数据结构基础掌握的不太行，需要恶补一下

func main() {
	list1 := &ListNode{1, nil}
	list1.Add(4)
	list1.Add(5)
	list2 := &ListNode{1, nil}
	list2.Add(3)
	list2.Add(4)
	list3 := &ListNode{2, nil}
	list3.Add(6)
	lists := []*ListNode{list1, list2, list3}
	mergeKLists(lists).traversal()
}
