package main

import (
	"fmt"
)

/*
Problem:
Given a 32-bit signed integer, reverse digits of an integer.

Example 1:

Input: 123
Output: 321
Example 2:

Input: -123
Output: -321
Example 3:

Input: 120
Output: 21
Note:
Assume we are dealing with an environment which could only store integers within the 32-bit signed integer range: [−2^31,  2^31 − 1].
For the purpose of this problem, assume that your function returns 0 when the reversed integer overflows.

翻译：
给你一个32位有符号整数，把这个数字翻转
假设我们只需要处理32位整数，范围 [−2^31,  2^31 − 1]，根据这个假设，如果反转后整数溢出，那么返回0
*/

// 解法一：运用标准库的函数
/*Runtime: 4 ms, faster than 100.00% of Go online submissions for Reverse Integer.
Memory Usage: 2.1 MB, less than 100.00% of Go online submissions for Reverse Integer.*/
/*
func reverse(x int) int {
	xStr := strconv.Itoa(x)
	var newStr []byte
	if xStr[0] == '-' {
		newStr = append(newStr, '-')
		for i := len(xStr) - 1; i >= 1; i-- {
			newStr = append(newStr, xStr[i])
		}
	} else {
		for i := len(xStr) - 1; i >= 0; i-- {
			newStr = append(newStr, xStr[i])
		}
	}
	resNum, _ := strconv.Atoi(string(newStr))
	if resNum > (2<<30-1) || resNum < (-2<<30) {
		return 0
	}
	return resNum
}*/

// 解法二：用取模的方法不断取末位数
// 时间复杂度 需要翻转的整数的位数为 log_10 n,故时间复杂度为O(log(n)) 空间复杂度 O(1)
func reverse(x int) int {
	resNum := 0
	for x != 0 {
		last := x % 10
		x /= 10
		resNum = resNum*10 + last
	}
	if resNum > (2<<30-1) || resNum < (-2<<30) {
		return 0
	}

	return resNum

}
func main() {
	ex1 := -123
	// fmt.Println(strconv.Itoa(ex1))
	fmt.Println(reverse(ex1))
}
