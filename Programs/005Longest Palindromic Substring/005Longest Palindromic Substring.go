package main

import "fmt"

/*
Problem:
Given a string s, find the longest palindromic substring in s. You may assume that the maximum length of s is 1000.

Example 1:

Input: "babad"
Output: "bab"
Note: "aba" is also a valid answer.
Example 2:

Input: "cbbd"
Output: "bb"

翻译：
给你一个字符串s，找到最长的回文子串
假定s的最大长度是1000

*/

func isLongestPalindrome(s string) bool {
	strLen := len(s)
	for i := 0; i < strLen/2; i++ {
		if s[i] != s[strLen-i-1] {
			return false
		}
	}
	return true
}

// 解法一：穷举法，写一个函数用于判断是否是回文子串。遍历找到所有子串，一一判断替换最大值
// 注意一点 当找到一个回文子串后， 必须找到比他更大的子串，所以不需要遍历长度更小的范围了。
// 时间复杂度 O(n^3)除了找子串的两层循环，验证也是一层循环 空间复杂度O(1)
/*
func longestPalindrome(s string) string {
	var theLongestPalindrome string
	longestLen := 0
	sLen := len(s)
	for i := 0; i < sLen; i++ {
		// 当剩余长度小于现有的最长子串，不需要继续找下去了
		if sLen-i <= longestLen {
			break
		}
		for j := i + longestLen; j < sLen+1; j++ {
			if isLongestPalindrome(s[i:j]) {
				if len(s[i:j]) > longestLen {
					longestLen = len(s[i:j])
					theLongestPalindrome = string(s[i:j])
				}
			}
		}
	}
	return theLongestPalindrome
}*/

// 解法二：中心扩散算法，由于回文的两侧互为镜像，可以以任意一点作为中心展开，可以知道总共有2n-1个中心
// 因为n为偶数的时候，中心可以处于两个字母中间
// 整体时间复杂度为 O(n^2)围绕中心找回文是一层循环，遍历中心又是一层循环
// 空间复杂度O(1)
/*func expandAroundCenter(s string, left int, right int) int {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left--
		right++
	}
	//退出循环前的时候是回文，所以长度按前一次的时候计算
	// 即 (right-1)-(left+1)+1
	return right - left - 1

}

func longestPalindrome(s string) string {
	sLen := len(s)
	if sLen < 1 {
		return ""
	}
	var start, end, maxLen int
	for i := 0; i < sLen; i++ {
		// 中心扩散求奇数回文最大长度
		len1 := expandAroundCenter(s, i, i)

		// 中心扩散求偶数回文最大长度
		len2 := expandAroundCenter(s, i, i+1)

		// 求出奇数回文偶数回文的最大长度值
		if len1 > len2 {
			maxLen = len1
		} else {
			maxLen = len2
		}

		// 根据长度算出回文的左边下标和右边下标
		if maxLen > end-start {
			// 若最大长度为奇数，那么一半的长度是 maxLen-1/2，但是奇数情况i是中心点，所以i-1-start+1= (maxLen-1)/2
			// 若为偶数，一半的长度是 maxLen/2,中心点在i和i+1中间，所以i-start+1= maxLen/2
			// 偶数的话会少减一个一，精度上讲是对的
			start = i - (maxLen-1)/2
			end = i + maxLen/2

		}
	}
	return s[start : end+1]

}*/

// 解法三：动态规划
// 时间复杂度和空间复杂度均为O(n^2)
/*如果我们已经知道“bab” 是回文，那么很明显，“ababa” 一定是回文，因为它的左首字母和右尾字母是相同的。
用二维数组记录每一维的数字分别记录回文串起点和终点位置
dp[j][i]   true 如果子串是回文串
dp[j][i]   false 子串不是回文串
即
dp[j][i] = (s[i] == s[j] && dp[j+1][i-1])
依据此递推式，我们先计算一字母的回文，再计算两字母。。。
*/
func longestPalindrome(s string) string {
	sLen := len(s)
	var dp [1000][1000]bool

	maxLen := 0 //保存最长回文子串长度
	start := 0  //保存最长回文子串起点
	for i := 0; i < sLen; i++ {
		for j := 0; j <= i; j++ {
			if i-j < 2 {
				dp[j][i] = (s[i] == s[j])
			} else {
				dp[j][i] = (s[i] == s[j] && dp[j+1][i-1])
			}
			if dp[j][i] && maxLen < i-j+1 {
				maxLen = i - j + 1
				start = j
			}
		}
	}

	return s[start : start+maxLen]
}
func main() {
	// fmt.Println(isLongestPalindrome("ababsa"))
	fmt.Println(longestPalindrome("ababsa"))
	// var dp [5][5]bool
	// fmt.Println(dp)
}
