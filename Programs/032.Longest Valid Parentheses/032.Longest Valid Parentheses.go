package main

import (
	"fmt"
)

/*
Given a string containing just the characters '(' and ')', find the length of the longest valid (well-formed) parentheses substring.

Example 1:

Input: "(()"
Output: 2
Explanation: The longest valid parentheses substring is "()"
Example 2:

Input: ")()())"
Output: 4
Explanation: The longest valid parentheses substring is "()()"


给定一个只包含 '(' 和 ')' 的字符串，找出最长的包含有效括号的子串的长度。

示例 1:

输入: "(()"
输出: 2
解释: 最长有效括号子串为 "()"
示例 2:

输入: ")()())"
输出: 4
解释: 最长有效括号子串为 "()()"
*/

// 之前做过一个判断括号是否有效的方法，可以截取字符串，用那个进行判断，明显效率并不高

// 解法一：可以这样操作，**用栈存放左括号的下标** 只有遇到左括号的时候放入栈中，
// 遇到右括号时，两种情况，如果栈为空，跳过，因为这个括号肯定不是有效括号队里的，此时更新起始点下标
// 如果栈不为空，那么栈顶元素出栈，因为都是左括号，计算当前右括号和有效括号起始点下标的距离，并更新最大值
// 出栈后，判断栈是否为空，如果为空，不能保证接下来的括号不能匹配，所以先判断此时的值，与最大值比较更新
// 如果不为空，同样是更新最大值，比较此时）下标和栈顶下标距离和当前最大值，进行更新
// 时间复杂度 O(n) 空间复杂度O(n)
/*
Runtime: 0 ms, faster than 100.00% of Go online submissions for Longest Valid Parentheses.
Memory Usage: 2.7 MB, less than 100.00% of Go online submissions for Longest Valid Parentheses.
*/
// import "./stack"
/*
因为用了自己实现的stack，不好放到leetcode上提交
func longestValidParentheses(s string) int {
	stackNew := stack.New()
	sLen := len(s)
	maxLength := 0
	start := 0
	for i := 0; i < sLen; i++ {
		if s[i] == '(' {
			stackNew.Push(i)
		} else if s[i] == ')' {
			if stackNew.Empty() {
				start = i
			} else {
				stackNew.Pop()
				maxLength = i - start
			}
		}
	}
	return maxLength
}*/

type Stack struct {
	list []int
	len  int
}

func (s *Stack) Push(v int) {
	s.list[s.len] = v
	s.len++
}

func (s *Stack) Empty() bool {
	if s.len == 0 {
		return true
	}
	return false
}
func (s *Stack) Pop() int {
	s.len--
	return s.list[s.len]
}
func (s *Stack) Top() int {
	if s.len >= 1 {
		return s.list[s.len-1]
	}
	return -1
}
func longestValidParentheses(s string) int {
	sLen := len(s)
	stackNew := new(Stack)
	stackNew.list = make([]int, sLen)
	stackNew.len = 0
	maxLength := 0
	start := 0
	for i := 0; i < sLen; i++ {
		// 如果是左括号，入栈
		if s[i] == '(' {
			stackNew.Push(i)
		} else  {
			// 如果是右括号，并且栈为空，跳过，因为这种情况一定不能匹配，此时需要更新有效括号其实下标，更新值为当前的下标+1
			if stackNew.Empty() {
				start = i+1
			} else {
				// 如果栈不为空，那么取出栈顶的值匹配，因为栈里都是左括号，所以肯定能匹配
				stackNew.Pop()
				// pop完后，两种情况，如果栈为空，那么和之前记录的有效括号起始下标比较，更新最大值
				if stackNew.Empty() {
					newLength := i - start+1
					if newLength > maxLength {
						maxLength = newLength
					}
				} else {
					// 如果pop后栈不为空，说明栈里面可能有多余的左括号，而且不能保证这个左括号一定能匹配，所以此时更新最大值
					// 这时更新最大值，要对比栈顶元素的下标。即用当前（的下标和栈顶元素下标的距离，和maxLength比较
					newLength := i - stackNew.Top()
					if newLength > maxLength {
						maxLength = newLength
					}
				}
			}
		}
	}
	return maxLength
}

/*
老外的实现比我的看起来简洁一些
func longestValidParentheses(s string) int {
	var stack []int
	max := 0
	stack = append(stack,-1)
	for i := 0; i < len(s);i ++ {
		if s[i]=='(' {
			stack = append(stack,i)
		}else {
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				stack = append(stack,i)
			}else {
				max = maxInt(max,i - stack[len(stack)-1])
			}
		}
	}
	return max
}

func maxInt(a,b int)int  {
	if a > b  {
		return a
	}
	return b
}
*/
func main() {
	// input := "(()"
	input := ")()())"
	fmt.Println(longestValidParentheses(input))
}
