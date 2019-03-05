package main

import (
	"fmt"
	"strconv"
	"strings"
)

/*
Problem:

Implement atoi which converts a string to an integer.

The function first discards as many whitespace characters as necessary until the first non-whitespace character is found. Then, starting from this character, takes an optional initial plus or minus sign followed by as many numerical digits as possible, and interprets them as a numerical value.

The string can contain additional characters after those that form the integral number, which are ignored and have no effect on the behavior of this function.

If the first sequence of non-whitespace characters in str is not a valid integral number, or if no such sequence exists because either str is empty or it contains only whitespace characters, no conversion is performed.

If no valid conversion could be performed, a zero value is returned.

Note:

Only the space character ' ' is considered as whitespace character.
Assume we are dealing with an environment which could only store integers within the 32-bit signed integer range: [−231,  231 − 1]. If the numerical value is out of the range of representable values, INT_MAX (231 − 1) or INT_MIN (−231) is returned.
Example 1:

Input: "42"
Output: 42
Example 2:

Input: "   -42"
Output: -42
Explanation: The first non-whitespace character is '-', which is the minus sign.
             Then take as many numerical digits as possible, which gets 42.
Example 3:

Input: "4193 with words"
Output: 4193
Explanation: Conversion stops at digit '3' as the next character is not a numerical digit.
Example 4:

Input: "words and 987"
Output: 0
Explanation: The first non-whitespace character is 'w', which is not a numerical
             digit or a +/- sign. Therefore no valid conversion could be performed.
Example 5:

Input: "-91283472332"
Output: -2147483648
Explanation: The number "-91283472332" is out of the range of a 32-bit signed integer.
             Thefore INT_MIN (−231) is returned.


翻译：
题目太长，直接搬leetcode中文站的翻译了

实现atoi方法，用于转化一个字符串为整数

首先，该函数会根据需要丢弃无用的开头空格字符，直到寻找到第一个非空格的字符为止。

当我们寻找到的第一个非空字符为正或者负号时，则将该符号与之后面尽可能多的连续数字组合起来，作为该整数的正负号；假如第一个非空字符是数字，则直接将其与之后连续的数字字符组合起来，形成整数。

该字符串除了有效的整数部分之后也可能会存在多余的字符，这些字符可以被忽略，它们对于函数不应该造成影响。

注意：假如该字符串中的第一个非空格字符不是一个有效整数字符、字符串为空或字符串仅包含空白字符时，则你的函数不需要进行转换。

在任何情况下，若函数不能进行有效的转换时，请返回 0。

说明：

假设我们的环境只能存储 32 位大小的有符号整数，那么其数值范围为 [−231,  231 − 1]。如果数值超过这个范围，qing返回  INT_MAX (231 − 1) 或 INT_MIN (−231) 。
*/

//题目比较没意思，go语言int在我这台上是64位，判断溢出没有必要，官方标准库已经很好用了，所以这题调用标准库
// 感觉go的标准库相对python还是难用多了，所以还是自己判断吧
// 主要问题在于判断各种情况比较麻烦
func myAtoi(str string) int {
	// 左边如果有空格妨碍判断，先去掉
	str = strings.TrimLeft(str, " ")
	// 如果是空串，那么直接返回0
	if str == "" {
		return 0
	}
	// 定义一个byte的切片，用于存放
	var numByte []byte
	// 去除空格后，如果首位是正负号，要加入byte切片中
	if str[0] == '-' || str[0] == '+' {
		numByte = append(numByte, str[0])
		// 后续碰到不是数字的就退出
		for i := 1; i < len(str); i++ {
			if str[i] >= '0' && str[i] <= '9' {
				numByte = append(numByte, str[i])
			} else {
				break
			}
		}

	} else {
		// 如果首位不是正负号，那么还是碰到不是数字的就退出
		for i := 0; i < len(str); i++ {
			if str[i] >= '0' && str[i] <= '9' {
				numByte = append(numByte, str[i])
			} else {
				break
			}
		}
	}
	resNum, _ := strconv.Atoi(string(numByte))
	if resNum < -2147483648 {
		return -2147483648
	} else if resNum > 2147483647 {
		return 2147483647
	}
	return resNum
}

func main() {
	input := " -42"
	input2 := "4193 with words"
	input3 := "+1"
	fmt.Println(myAtoi(input))
	fmt.Println(myAtoi(input2))
	fmt.Println(myAtoi(input3))
}
