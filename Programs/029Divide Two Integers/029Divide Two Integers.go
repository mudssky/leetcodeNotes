package main

import "fmt"

/*
Given two integers dividend and divisor, divide two integers without using multiplication, division and mod operator.

Return the quotient after dividing dividend by divisor.

The integer division should truncate toward zero.

Example 1:

Input: dividend = 10, divisor = 3
Output: 3
Example 2:

Input: dividend = 7, divisor = -3
Output: -2
Note:

Both dividend and divisor will be 32-bit signed integers.
The divisor will never be 0.
Assume we are dealing with an environment which could only store integers within the 32-bit signed integer range: [−231,  231 − 1]. For the purpose of this problem, assume that your function returns 231 − 1 when the division result overflows.


给定两个整数，被除数 dividend 和除数 divisor。将两数相除，要求不使用乘法、除法和 mod 运算符。

返回被除数 dividend 除以除数 divisor 得到的商。

示例 1:

输入: dividend = 10, divisor = 3
输出: 3
示例 2:

输入: dividend = 7, divisor = -3
输出: -2
说明:

被除数和除数均为 32 位有符号整数。
除数不为 0。
假设我们的环境只能存储 32 位有符号整数，其数值范围是 [−231,  231 − 1]。本题中，如果除法结果溢出，则返回 231 − 1。
*/

// 解法一：虽然不满足题目的要求，就使用正常除法
/*
Runtime: 4 ms, faster than 100.00% of Go online submissions for Divide Two Integers.
Memory Usage: 2.4 MB, less than 44.44% of Go online submissions for Divide Two Integers.
*/
/*func divide(dividend int, divisor int) int {
	res := dividend / divisor
	if res > 2<<30-1 || res < -2<<30 {
		return 2<<30 - 1
	}
	return res
}*/
// 解法二： 方案是用加减法代替除法。两个数都是正数的情况，即转化为：被除数，减去多少个除数以后结果开始小于除数
// 但是一个个减，循环的次数比较多，太慢了，我们可以用位移运算，加快速度，减少运算次数
// 已知divisor<<i = divisor * 2<<i-1,
// 我们可以遍历2的指数，即判断divisor<<i和被除数的大小，刚好比被除数小时，这部分个数的被除数即计算完毕
// 此时我们从被除数中减去divisor * 2<<i-1，结果中加上 2<<i-1
// 重复这种运算，直至，结果小于被除数
/*func divide(dividend int, divisor int) int {
	// 先判断两个数的正负，分4种情况，确定最终结果的正负
	// 把两个数都转为正数，方便计算
	flag := 1
	if dividend < 0 && divisor < 0 {
		dividend = -dividend
		divisor = -divisor
	} else if dividend < 0 && divisor > 0 {
		dividend = -dividend
		flag = -1
	} else if dividend > 0 && divisor < 0 {
		divisor = -divisor
		flag = -1
	}
	// 如果两个数都是正数，不用进行任何处理
	res := 0
	var i uint
	for dividend >= divisor {
		i = 0
		// fmt.Println(res)
		for ; i < 33; i++ {
			// fmt.Println(i)
			if divisor<<i > dividend {
				if i == 1 {
					res++
					dividend -= divisor
					break
				} else {

					// divisor<<(i-1)= divisor * 2^(i-1)=2<<(i-2)
					res += 2 << (i - 2)
					dividend -= divisor << (i - 1)
					break
				}
			}
		}
	}
	// res = res * flag
	// 分解成判断
	if flag == -1 {
		res = -res
	}
	if res > 2<<30-1 {
		return 2<<30 - 1
	}
	return res
}*/

// 解法二 优化
// 解法二的循环还可以优化一下，因为i每次从最小0开始遍历，有不少遍历重复了，如果一开始从最大的开始遍历就减少了重复的遍历
func divide(dividend int, divisor int) int {
	// 先判断两个数的正负，分4种情况，确定最终结果的正负
	// 把两个数都转为正数，方便计算
	flag := 1
	if dividend < 0 && divisor < 0 {
		dividend = -dividend
		divisor = -divisor
	} else if dividend < 0 && divisor > 0 {
		dividend = -dividend
		flag = -1
	} else if dividend > 0 && divisor < 0 {
		divisor = -divisor
		flag = -1
	}
	// 如果两个数都是正数，不用进行任何处理
	res := 0
	// go语言中int默认64位 这里位移值31刚好是边界值，32的时候，位移会出现溢出，导致程序不符合预期
	var i uint = 31
	for dividend >= divisor {
		for ; i >= 0; i-- {
			if divisor<<i <= dividend {
				// divisor<<i= divisor * 2^(i-1)=2<<(i-1)
				if i == 0 {
					res++
					dividend -= divisor
					i--
					break
				} else {
					res += 2 << (i - 1)
					dividend -= divisor << i
					i--
					break
				}
			}
		}
	}
	// res = res * flag
	// 分解成判断
	if flag == -1 {
		res = -res
	}
	fmt.Println(res)
	if res > 2<<30-1 {
		return 2<<30 - 1
	}
	return res
}

func main() {
	fmt.Println(divide(-2147483648, -2147483648))

}
