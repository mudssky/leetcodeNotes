package main

import (
	"fmt"
)

/*
Problem:
Given a string, find the length of the longest substring without repeating characters.

Example 1:

Input: "abcabcbb"
Output: 3
Explanation: The answer is "abc", with the length of 3.
Example 2:

Input: "bbbbb"
Output: 1
Explanation: The answer is "b", with the length of 1.
Example 3:

Input: "pwwkew"
Output: 3
Explanation: The answer is "wke", with the length of 3.
             Note that the answer must be a substring, "pwke" is a subsequence and not a substring.

翻译：

给你一个字符串，找到最长子串求出长度，不能有重复的字母

*/
// 解法一：暴力求解，遍历所有组合，用一个map判断字母是否重复，提高效率，不用map的话还要再遍历一层，相当于三层循环
//时间复杂度 O(n^2) 空间复杂度 因为要用一个字典，所以最多是把所有元素放进去,相当于字符串的长度O(n)
/*
Runtime: 356 ms, faster than 12.86% of Go online submissions for Longest Substring Without Repeating Characters.
Memory Usage: 6.9 MB, less than 5.30% of Go online submissions for Longest Substring Without Repeating Characters.
*/
/*func lengthOfLongestSubstring(s string) int {

	var max int
	for i := 0; i < len(s); i++ {
		dict := make(map[byte]int)
		for j := i; j < len(s); j++ {
			if _, ok := dict[s[j]]; ok {
				if j-i > max {
					max = j - i
				}
				break
			}
			if j == len(s)-1 {
				if max < len(s)-i {
					max = len(s) - i
					break
				}

			}
			dict[s[j]] = j
		}
	}
	return max
}*/
// 优化，当剩余的长度小于或等于max的时候，可以不用继续遍历了
// 把len(s)这类重复使用多次的数值，用一个变量存下来
/*
Runtime: 348 ms, faster than 13.24% of Go online submissions for Longest Substring Without Repeating Characters.
Memory Usage: 6.9 MB, less than 5.30% of Go online submissions for Longest Substring Without Repeating Characters.
*/
/*
func lengthOfLongestSubstring(s string) int {

	var max int
	strLen := len(s)
	for i := 0; i < strLen; i++ {
		// 当剩余的长度小于或等于max的时候，可以不用继续遍历了
		if strLen-i <= max {
			break
		}
		dict := make(map[byte]int)
		for j := i; j < strLen; j++ {
			if _, ok := dict[s[j]]; ok {
				if j-i > max {
					max = j - i
				}
				break
			}
			if j == strLen-1 {
				left := strLen - i
				if max < left {
					max = left
					break
				}

			}
			dict[s[j]] = j
		}
	}
	return max
}*/

// 解法二：实际上解法一进行了很多重复的计算，比如说知道si-sj之间没有重复的字符以后，这一段就不用再判断了
// 我们可以维护一个切片的范围，出现了重复的字段，就把左边的下标移到被重复的字段之后
// 重复时这个切片的范围的长度大于已经找到的最大值，那么就把最大值更新
//用map记录已经存在的值，后面的值会把前面的覆盖
/*
Runtime: 12 ms, faster than 69.67% of Go online submissions for Longest Substring Without Repeating Characters.
Memory Usage: 3 MB, less than 63.25% of Go online submissions for Longest Substring Without Repeating Characters.
时间复杂度：O(n) 空间复杂度O(n)
*/
/*
func lengthOfLongestSubstring(s string) int {
	var max int
	dict := make(map[byte]int)
	start := 0
	strLen := len(s)
	for i := 0; i < strLen; i++ {
		//当剩余的长度小于或等于max的时候，可以不用继续遍历了
		if strLen-start <= max {
			break
		}

		//判断某个字符在字典中是否存在，即使存在还必须满足这个找到的下标在维护的切片范围之内
		//条件满足的情况下，说明出现了重复，把这个范围的长度和最大值比较决定是否更新
		// 同时移动左边边界的下标到最初的重复值之后
		// 出现重复的情况，未重复的部分是范围的长度-1，也就是不计算最右端的
		if _, ok := dict[s[i]]; ok && dict[s[i]] >= start {
			if i-start > max {
				max = i - start
			}
			start = dict[s[i]] + 1
		}
		dict[s[i]] = i
		// 因为上面的判断没有计算维护的切片最右端的元素，如果到最后一个单位，不会被计数进去，
		//所以当遍历到最后的时候，判断一下
		if strLen-1 == i {
			if i-start+1 > max {
				max = i - start + 1
			}
		}
	}

	return max
}
*/
//解法三：如果字符集为ascii码，那么我们可以用一个[128]int{}的数组来作为一个map使用。
func lengthOfLongestSubstring(s string) int {
	var max int
	dict := [128]int{}
	// go语言数组初始赋值为-1，因为0也在字符集中，所以不适合做初始值
	for i := 0; i < 128; i++ {
		dict[i] = -1
	}
	start := 0
	for i, v := range s {
		if dict[v] != -1 && dict[v] >= start {
			if max < i-start {
				max = i - start
			}
			start = dict[v] + 1
		}
		dict[v] = i
		if i == len(s)-1 {
			if i-start+1 > max {
				max = i - start + 1
			}
		}
	}
	return max
}
func main() {
	s1 := "abcabcbb" //3
	s2 := "aab"      //2
	s3 := "abba"     //2
	s4 := "bbbbb"    //1
	s5 := "tmmzuxt"  //5
	fmt.Println(lengthOfLongestSubstring(s1))
	fmt.Println(lengthOfLongestSubstring(s2))
	fmt.Println(lengthOfLongestSubstring(s3))
	fmt.Println(lengthOfLongestSubstring(s4))
	fmt.Println(lengthOfLongestSubstring(s5))
}
