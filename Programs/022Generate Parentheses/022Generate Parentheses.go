package main

import "fmt"

/*
Given n pairs of parentheses, write a function to generate all combinations of well-formed parentheses.

For example, given n = 3, a solution set is:

[
  "((()))",
  "(()())",
  "(())()",
  "()(())",
  "()()()"
]

给出 n 代表生成括号的对数，请你写出一个函数，使其能够生成所有可能的并且有效的括号组合。

例如，给出 n = 3，生成结果为：

[
  "((()))",
  "(()())",
  "(())()",
  "()(())",
  "()()()"
]
*/

// 解法一：暴力求解计算2n个字符，递归查找每个分别是左括号或者右的所有可能，判断符合要求的留下来。
// 这样做要计算大量不符合要求的字符串，浪费很多资源
// 时间复杂度O(n* 2^2n),空间复杂度也是指数级的。

func isValid(s string) bool {
	stack := make([]byte, len(s))
	stack[0] = s[0]
	top := 1
	for i := 1; i < len(s); i++ {
		if top >= 1 && stack[top-1] == '(' && s[i] == ')' {
			top--
		} else {
			stack[top] = s[i]
			top++
		}
	}
	if top == 0 {
		return true
	}
	return false
}

/*
func generateAll(s string, n int, res *[]string, index int) {
	if index == n {
		if isValid(s) {
			*res = append(*res, s)
			// fmt.Println(s)
		}
		return
	}
	generateAll(s+"(", n, res, index+1)
	generateAll(s+")", n, res, index+1)
}

func generateParenthesis(n int) []string {
	res := []string{}
	generateAll("", 2*n, &res, 0)
	return res
}
*/
// 优化解法一，其实左括号或者右括号的数量超过一半之后，已经不用判断了。所以我们可以记录一下左括号的数目
// 还有右括号数量大于或等于括号时，这时候也不应该往后加右括号

func generateAll(s string, n int, res *[]string, left int, right int) {
	if len(s) == 2*n {
		if isValid(s) {
			*res = append(*res, s)
		}
		return
	}
	if left < n {
		generateAll(s+"(", n, res, left+1, right)
	}
	if right < left {
		generateAll(s+")", n, res, left, right+1)
	}
}

func generateParenthesis(n int) []string {
	res := []string{}
	generateAll("", n, &res, 0, 0)
	return res
}

func main() {
	fmt.Println(generateParenthesis(3))
}
