package main

import "fmt"

/*
Problem:
You are given two non-empty linked lists representing two non-negative integers. The digits are stored in reverse order and each of their nodes contain a single digit. Add the two numbers and return it as a linked list.

You may assume the two numbers do not contain any leading zero, except the number 0 itself.

Example:

Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
Output: 7 -> 0 -> 8
Explanation: 342 + 465 = 807.

翻译：
给你两个非空链表代表两个非负整数，每一位数字都都是按照逆序存储的，把这两个数相加，并且用链表形式返回
假定两个数开头都不包含0，除了他们自己是0的时候


*/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

//ListNode 单向链表
type ListNode struct {
	Val  int
	Next *ListNode
}

func (ln *ListNode) traversal() {
	for {
		if ln.Next != nil {
			fmt.Print(ln.Val, ',')
		} else {
			fmt.Print(ln.Val, '\n')
			break
		}
		ln = ln.Next
	}
}

//Add 单向列表添加元素在最后面
func (ln *ListNode) Add(val int) {
	// 初始值明显不是nil，本来想写个判断，分辨空节点的情况
	// 不过好像不好分辨，只能干脆选择默认添加到最后面的方式
	/*
		if ln == nil {
			ln = &ListNode{val, nil}
		} else {
			for ln.Next != nil {
				ln = ln.Next
			}
			ln.Next = &ListNode{val, nil}
		}*/
	for ln.Next != nil {
		ln = ln.Next
	}
	ln.Next = &ListNode{val, nil}

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

//解法一：小学算数，模拟我们学过的加法运算，按位遍历，每位执行加法，并且用一个标志位表示进位
// 时间复杂度O(max(m,n)) 空间复杂度O(max(m,n))
/*
执行用时: 52 ms, 在Add Two Numbers的Go提交中击败了5.28% 的用户
内存消耗: 5.2 MB, 在Add Two Numbers的Go提交中击败了9.88% 的用户
这结果有点垃圾，明显需要优化
*/
/*
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	l3 := &ListNode{}
	startNode := l3
	flag := 0
	for {

		// fmt.Println(l1.Val, l2.Val, flag)
		res := l1.Val + l2.Val + flag
		if res > 9 {
			l3.Val = res - 10
			//flag用来记录上次计算的进位
			flag = 1
		} else {
			l3.Val = res
			flag = 0
		}
		//判断放的位置很重要，应该在最后一个链表节点数据计算之后
		if l1.Next == nil && l2.Next == nil {
			//最后一次计算完还有进位没有判断
			if flag == 1 {
				l3.Next = &ListNode{1, nil}
			}
			return startNode
		}
		l3.Next = &ListNode{}
		l3 = l3.Next
		//如果一个链表到头了，那么之后位数的数字都用0来加，这里有一点不好就是改变了原来链表的值，如果两个链表长度不一样的话，有一个的值会被改变
		// 可以分别用一个标志位来记录两个链表到没到头，然后前面加上判断,逻辑就变得相当复杂了,所以就不改了
		if l1.Next != nil {
			l1 = l1.Next
		} else {
			l1.Val = 0
		}
		if l2.Next != nil {
			l2 = l2.Next
		} else {
			l2.Val = 0
		}

	}
}*/

//操作优化
/*
Runtime: 12 ms, faster than 100.00% of Go online submissions for Add Two Numbers.
Memory Usage: 5 MB, less than 83.89% of Go online submissions for Add Two Numbers.
*/
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	l3 := &ListNode{}
	startNode := l3
	var flag int
	var sum int
	for {
		// 原先的判断和多余的赋值用取余和计算完成
		sum = l1.Val + l2.Val + flag
		l3.Val = sum % 10
		flag = sum / 10
		//判断放的位置很重要，应该在最后一个链表节点数据计算之后
		if l1.Next == nil && l2.Next == nil {
			//最后一次计算完还有进位没有判断
			if flag == 1 {
				l3.Next = &ListNode{1, nil}
			}
			return startNode
		}
		l3.Next = &ListNode{}
		l3 = l3.Next
		// 下一个节点是nil说明已经到头了
		if l1.Next == nil {
			l1.Val = 0
		} else {
			l1 = l1.Next
		}
		if l2.Next == nil {
			l2.Val = 0
		} else {
			l2 = l2.Next
		}
	}
}

// 解法二：运用递归，先加完两个链表的第一位，那么结果就是以第一位的结果创建节点，连接去除第一位的链表相加的结果
// 退出条件就是两个链表都走到尾部，如果只有一个走到尾部，那就在那个链表尾部接值为0节点续上
/*func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	//计算第一位相加的结果
	val := l1.Val + l2.Val
	// 如果大于9，将进位转移到两个随便一个链表的下一位
	if val >= 10 {
		val = val % 10
		if l1.Next == nil {
			l1.Next = &ListNode{0, nil}
		}
		if l2.Next == nil {
			l2.Next = &ListNode{0, nil}
		}
		l1.Next.Val++
	}
	//退出条件，如果两个链表都走到尾部，就退出
	if l1.Next == nil && l2.Next == nil {
		return &ListNode{val, nil}
	} else {
		//如果没有都走到尾部，那么判断只有一个走到尾部的情况把尾部续一个值为0的节点，和上面>9的逻辑一样
		if l1.Next == nil {
			l1.Next = &ListNode{0, nil}
		}
		if l2.Next == nil {
			l2.Next = &ListNode{0, nil}
		}
	}
	// 递归继续计算剩下的位数的结果，当前位的结果为val
	res := addTwoNumbers(l1.Next, l2.Next)
	return &ListNode{val, res}
}*/

func main() {
	l1 := new(ListNode)
	l2 := new(ListNode)
	l1.Add(2)
	l1.Add(4)
	l1.Add(3)
	l2.Add(5)
	l2.Add(6)
	l2.Add(4)
	// l1 := &ListNode{2, &ListNode{4, &ListNode{3, nil}}}
	// l2 := &ListNode{5, &ListNode{6, &ListNode{4, nil}}}
	traversal(l1)
	traversal(l2)
	startNode := addTwoNumbers(l1, l2)
	// startNode.traversal()
	traversal(startNode)
	traversal(l1)
	traversal(l2)
}
