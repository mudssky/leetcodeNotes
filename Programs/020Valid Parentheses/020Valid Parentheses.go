package main

/*
Given a string containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.

An input string is valid if:

Open brackets must be closed by the same type of brackets.
Open brackets must be closed in the correct order.
Note that an empty string is also considered valid.

Example 1:

Input: "()"
Output: true
Example 2:

Input: "()[]{}"
Output: true
Example 3:

Input: "(]"
Output: false
Example 4:

Input: "([)]"
Output: false
Example 5:

Input: "{[]}"
Output: true
给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效。

有效字符串需满足：

左括号必须用相同类型的右括号闭合。
左括号必须以正确的顺序闭合。
注意空字符串可被认为是有效字符串。

示例 1:

输入: "()"
输出: true
示例 2:

输入: "()[]{}"
输出: true
示例 3:

输入: "(]"
输出: false
示例 4:

输入: "([)]"
输出: false
示例 5:

输入: "{[]}"
输出: true
*/
import "fmt"

// 解法一使用压栈的方法，产生顺序符合并且匹配的两个括号就弹出，最后如果栈是空的说明所有括号都被匹配了。
// 时间复杂度 O(n),空间复杂度O(n)
/*Runtime: 0 ms, faster than 100.00% of Go online submissions for Valid Parentheses.
Memory Usage: 1.9 MB, less than 100.00% of Go online submissions for Valid Parentheses.*/
/*func bracketMatch(left, right byte) bool {
	if (left == '(' && right == ')') || (left == '[' && right == ']') || (left == '{' && right == '}') {
		return true
	}

	return false

}

func isValid(s string) bool {
	sLen := len(s)
	if sLen%2 != 0 {
		return false
	}
	if sLen == 0 {
		return true
	}
	stack := make([]byte, len(s))
	stack[0] = s[0]
	top := 1
	for i := 1; i < sLen; i++ {
		// 如果没有达到栈底进行这个判断，判断是否匹配
		if top >= 1 && bracketMatch(stack[top-1], s[i]) {
			top--
		} else {
			stack[top] = s[i]
			top++
		}
		fmt.Println(stack)
	}
	if top == 0 {
		return true
	}
	return false
}*/
// 解法一优化
// 判断是否匹配的一部写了几个||的逻辑，实际上可以用map代替，只需判断一次即可

/*时间复杂度：O(n)O(n)，因为我们一次只遍历给定的字符串中的一个字符并在栈上进行 O(1)O(1) 的推入和弹出操作。
空间复杂度：O(n)O(n)，当我们将所有的开括号都推到栈上时以及在最糟糕的情况下，我们最终要把所有括号推到栈上。例如 ((((((((((*/
/*
执行用时 : 0 ms, 在Valid Parentheses的Go提交中击败了100.00% 的用户
内存消耗 : 2 MB, 在Valid Parentheses的Go提交中击败了17.00% 的用户
*/
func bracketMatch(left, right byte) bool {
	dict := map[byte]byte{
		'{': '}',
		'[': ']',
		'(': ')',
	}
	if val, ok := dict[left]; ok && val == right {
		return true
	}
	return false
}
func isValid(s string) bool {

	sLen := len(s)
	// 若输入空字符串，有效
	if sLen == 0 {
		return true
	}
	// 若输入括号的数目是奇数，不可能完全匹配
	if sLen%2 != 0 {
		return false
	}
	// 用一个map存储映射关系
	dict := map[byte]byte{
		'{': '}',
		'[': ']',
		'(': ')',
	}
	// 用byte的切片模拟一个栈
	stack := make([]byte, len(s))
	stack[0] = s[0]
	top := 1
	for i := 1; i < sLen; i++ {
		// 未到达栈底时，可以进行判断
		if top >= 1 {
			if val, ok := dict[stack[top-1]]; ok && val == s[i] {
				top--
			} else {
				// 如果不匹配,继续添加这个byte到栈顶
				stack[top] = s[i]
				top++
			}
			// 若已经在栈底，添加这个byte到栈顶
		} else {
			stack[top] = s[i]
			top++
		}
		// fmt.Println(stack)
	}
	// 若最后栈为空,说明全部匹配完成
	if top == 0 {
		return true
	}
	return false
}

func main() {
	// input1:="{}"
	// input2 := "()[]{}"
	input3 := "{{)}"

	fmt.Println(isValid(input3))
}
