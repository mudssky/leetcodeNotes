package main

/*
Implement strStr().

Return the index of the first occurrence of needle in haystack, or -1 if needle is not part of haystack.

Example 1:

Input: haystack = "hello", needle = "ll"
Output: 2
Example 2:

Input: haystack = "aaaaa", needle = "bba"
Output: -1
Clarification:

What should we return when needle is an empty string? This is a great question to ask during an interview.

For the purpose of this problem, we will return 0 when needle is an empty string. This is consistent to C's strstr() and Java's indexOf().

实现 strStr() 函数。

给定一个 haystack 字符串和一个 needle 字符串，在 haystack 字符串中找出 needle 字符串出现的第一个位置 (从0开始)。如果不存在，则返回  -1。

示例 1:

输入: haystack = "hello", needle = "ll"
输出: 2
示例 2:

输入: haystack = "aaaaa", needle = "bba"
输出: -1
说明:

当 needle 是空字符串时，我们应当返回什么值呢？这是一个在面试中很好的问题。

对于本题而言，当 needle 是空字符串时我们应当返回 0 。这与C语言的 strstr() 以及 Java的 indexOf() 定义相符。
*/

// 解法一 此题不难，遍历一遍用切片截取等长字符串比较即可
/*
Runtime: 0 ms, faster than 100.00% of Go online submissions for Implement strStr().
Memory Usage: 2.2 MB, less than 100.00% of Go online submissions for Implement strStr().
*/

func strStr(haystack string, needle string) int {
	needleLen := len(needle)
	if needleLen == 0 {
		return 0
	}
	haystackLen := len(haystack)
	/*后面的循环已经可以判断这种情况
	if needleLen > haystackLen {
		return -1
	}*/
	for i := 0; i+needleLen <= haystackLen; i++ {
		if haystack[i:i+needleLen] == needle {
			return i
		}
	}
	return -1
}
func main() {

}
