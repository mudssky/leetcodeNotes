package main

import (
	"fmt"
	"regexp"
)

/*
Problem:
Given an input string (s) and a pattern (p), implement regular expression matching with support for '.' and '*'.

'.' Matches any single character.
'*' Matches zero or more of the preceding element.
The matching should cover the entire input string (not partial).

Note:

s could be empty and contains only lowercase letters a-z.
p could be empty and contains only lowercase letters a-z, and characters like . or *.
Example 1:

Input:
s = "aa"
p = "a"
Output: false
Explanation: "a" does not match the entire string "aa".
Example 2:

Input:
s = "aa"
p = "a*"
Output: true
Explanation: '*' means zero or more of the precedeng element, 'a'. Therefore, by repeating 'a' once, it becomes "aa".
Example 3:

Input:
s = "ab"
p = ".*"
Output: true
Explanation: ".*" means "zero or more (*) of any character (.)".
Example 4:

Input:
s = "aab"
p = "c*a*b"
Output: true
Explanation: c can be repeated 0 times, a can be repeated 1 time. Therefore it matches "aab".
Example 5:

Input:
s = "mississippi"
p = "mis*is*p*."
Output: false

翻译：
输入一个字符串s，和一个模式p，实现正则表达式支持'.'和'*'的匹配
'.'匹配单个字母
'*'匹配0个或多个前面的元素
匹配应该覆盖整个字符串，而不是部分字符串

说明：
s 可能为空，且只包含从 a-z 的小写字母。
p 可能为空，且只包含从 a-z 的小写字母，以及字符 . 和 *。

*/
// 解法一：调用标准库的正则
/*
Runtime: 4 ms, faster than 83.47% of Go online submissions for Regular Expression Matching.
Memory Usage: 6.6 MB, less than 10.71% of Go online submissions for Regular Expression Matching.
*/
/*
func isMatch(s string, p string) bool {
	p = "^" + p + "$"
	res, _ := regexp.Match(p, []byte(s))
	return res
}*/
// 解法二：递归求解
// 两种特殊字符.和*,如果没有*,判断就会容易很多,所以分两种情况，有*和没*，还要讨论特殊情况
// 一步步匹配模式字符串p
/*
Runtime: 12 ms, faster than 47.46% of Go online submissions for Regular Expression Matching.
Memory Usage: 2.1 MB, less than 100.00% of Go online submissions for Regular Expression Matching.
*/
/*func isMatch(s string, p string) bool {
	// 退出条件
	// 当匹配字符串长度是0的时候，也就是匹配完了，此时s如果正好匹配完就是匹配成功
	// 如果s还有剩余就是匹配失败
	if len(p) == 0 {
		if len(s) == 0 {
			return true
		} else {
			return false
		}
	}

	// 匹配第一个字符如果字符串为空,还有一种机会可以匹配成功,那就是后面有*
	var firstMarthed bool
	if s != "" && (p[0] == s[0] || p[0] == '.') {
		firstMarthed = true
	} else {
		firstMarthed = false
	}
	// 匹配完第一个字符,需要判断第二个字符是不是*
	if len(p) >= 2 && p[1] == '*' {
		// 如果后面有*字符,两种情况,一是前面的字符没匹配上,那么跳过这个x*的匹配继续往*后面匹配,因为*可以表示前面的字符出现0次.
		// 另一种情况是前面的字符匹配上了,此时我们仍然用这个模式匹配字符串的后面
		return isMatch(s, p[2:]) || (firstMarthed && isMatch(s[1:], p))
	} else {
		// 如果第一个字符匹配失败,而且后面没*,那么会因为短路运算符,直接返回失败
		// 如果匹配成功,并且后面没有*的情况下,自然是继续向下匹配
		return firstMarthed && isMatch(s[1:], p[1:])
	}
}*/
// 解法三: top-down 自顶向下法
// 对解法二可以进行优化,因为 isMatch(s[1:], p[1:])这些递归调用的步骤,有很多重复的计算
// 可以用map把这些计算存起来,减少多余的计算
// dp(i,j)代表问题 s[i:]和p[j:]是否匹配？
func isMatch(s string, p string) bool {
	// go语言大部分数据类型都是值类型，除了指针映射。。。
	// 结构体默认也是传值，可以作为map的键
	// 创建一个结构体作为map的键，dpMap可以存对应dp(i,j)的结果
	type ij struct {
		i int
		j int
	}
	dpMap := make(map[ij]bool)

	var ans bool
	// go语言在函数中声明函数是这样的方式
	var dp func(i, j int) bool
	dp = func(i, j int) bool {
		// 进入函数后，先搜索map看是否已经得到答案
		if _, ok := dpMap[ij{i, j}]; !ok {
			// 如果没有答案，那么还是原来的逻辑
			// 起初我们从最左端开始，也就是i和j分别是0
			// 退出条件是当p模式已经匹配到末尾，此时若s也到末尾返回true，否则结果为false
			if j == len(p) {
				ans = i == len(s)
			} else {

				var firstMarthed bool
				// 匹配第一个，如果i没到结尾，并且。p成功匹配第一个，那么第一个匹配成功
				// i到结尾，p没到结尾，那么只有一种p中含*的情况需要在后续匹配，这里我们先设置为false
				if i < len(s) && (p[j] == s[i] || p[j] == '.') {
					firstMarthed = true
				} else {
					firstMarthed = false
				}
				// 判断后一位包含*的情况
				// 若包含*，两种匹配情况，一是跳过这个*，二是在第一个匹配成功的情况下，继续向后匹配。
				if j+1 < len(p) && p[j+1] == '*' {
					ans = dp(i, j+2) || (firstMarthed && dp(i+1, j))
				} else {
					// 如果不含*，且第一个已经匹配成功，继续往下匹配
					ans = firstMarthed && dp(i+1, j+1)
				}

			}
			// 上面匹配玩会得到一个结果，把这个结果存到dpMap中
			dpMap[ij{i, j}] = ans

		}
		// map中能找到结果的情况下才会运行到这里，直接返回结果
		return dpMap[ij{i, j}]

	}
	return dp(0, 0)

}

func main() {

	fmt.Println(isMatch("a", "ab*a"))
	fmt.Println(isMatch("aa", "a"))
	test1, _ := regexp.Match("aa", []byte("a"))
	fmt.Println(test1)
	fmt.Println(isMatch("aa", "a"))
}
