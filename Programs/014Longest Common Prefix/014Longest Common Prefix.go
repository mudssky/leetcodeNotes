package main

import "fmt"

/*
Write a function to find the longest common prefix string amongst an array of strings.

If there is no common prefix, return an empty string "".

Example 1:

Input: ["flower","flow","flight"]
Output: "fl"
Example 2:

Input: ["dog","racecar","car"]
Output: ""
Explanation: There is no common prefix among the input strings.
Note:

All given inputs are in lowercase letters a-z.

翻译：
编写一个函数来查找字符串数组中的最长公共前缀。

如果不存在公共前缀，返回空字符串 ""。
*/
// 解法一： 先求出所有字符串中最短长度，然后慢慢缩短长度用切片进行比较
// 时间复杂度
// 最坏的情况下是 O(strlen*minlen)即遍历完都找不到字符串的情况，大约O(n^2)吧
// 空间复杂度，其中用到一个数组 O(minLen)
/*
Runtime: 0 ms, faster than 100.00% of Go online submissions for Longest Common Prefix.
Memory Usage: 2.3 MB, less than 100.00% of Go online submissions for Longest Common Prefix.
*/
func longestCommonPrefix(strs []string) string {
	strsLen := len(strs)
	if strsLen == 0 {
		return ""
	}
	minLen := len(strs[0])
	// 先遍历一边求出最短长度
	for i := 0; i < strsLen; i++ {
		if minLen > len(strs[i]) {
			minLen = len(strs[i])
		}
	}
	if minLen == 0 {
		return ""
	}
	// 取一个作为参照值
	dictStr := strs[0][:minLen]
	// 接下来从最长往最短遍历,用切片进行比较，减少循环层数
	for minLen > 0 {
		allHas := true
		for i := 0; i < strsLen; i++ {
			if strs[i][:minLen] != dictStr[:minLen] {
				allHas = false
			}
		}
		// 循环，到切片出现匹配位置，否则切片的末端不断往左移
		if allHas == true {
			return dictStr[:minLen]
		}
		minLen--
	}
	return ""
}
func main() {
	input := []string{"flower", "flow", "flight"}

	fmt.Println(longestCommonPrefix(input))
}
