package main

import (
	"fmt"
)

/*
The count-and-say sequence is the sequence of integers with the first five terms as following:

1.     1
2.     11
3.     21
4.     1211
5.     111221
1 is read off as "one 1" or 11.
11 is read off as "two 1s" or 21.
21 is read off as "one 2, then one 1" or 1211.

Given an integer n where 1 ≤ n ≤ 30, generate the nth term of the count-and-say sequence.

Note: Each term of the sequence of integers will be represented as a string.



Example 1:

Input: 1
Output: "1"
Example 2:

Input: 4
Output: "1211"

报数序列是一个整数序列，按照其中的整数的顺序进行报数，得到下一个数。其前五项如下：

1.     1
2.     11
3.     21
4.     1211
5.     111221
1 被读作  "one 1"  ("一个一") , 即 11。
11 被读作 "two 1s" ("两个一"）, 即 21。
21 被读作 "one 2",  "one 1" （"一个二" ,  "一个一") , 即 1211。

给定一个正整数 n（1 ≤ n ≤ 30），输出报数序列的第 n 项。

注意：整数顺序将表示为一个字符串。

分析，其实后面的数就是读前面的数产生的数
第一个是 ： 1
第二个是 读第一个， 一个一 也就是 11
第三个数 是读第二个 两个一， 也就是21
第四个再读第三个 一个二和一个一，也就是1211
即重复次数，加重复的数字
所以这题明显比较适合用递归来做
因为后面的一个数是数前面的数的结果。
*/

// 解法一：递归求解
// 时间复杂度 大于O(n)
// 做一个根据之前数字的序列，求当前序列的函数
/*func countBefore(s string) string {
	resStr := ""
	// 本质上是对每个数字用重复次数+重复的数字然后拼起来
	repeatCount := 1
	for i := 0; i < len(s); i++ {
		if i+1 < len(s) && s[i] == s[i+1] {
			// 寻找出现重复的数字，累计重复次数
			// 如果运行到末尾，且末尾不是重复数字的情况也是执行else的逻辑
			// 如果末尾和之前的重复，那么会被repeatCount记录
			repeatCount++
		} else {
			resStr += strconv.Itoa(repeatCount) + string(s[i])
			// 重置重复次数
			repeatCount = 1
		}
	}
	return resStr
}

func countAndSay(n int) string {
	if n == 1 {
		return "1"
	}
	return countBefore(countAndSay(n - 1))
}
*/

// 递归消耗内存太多了,拆成普通的循环，同时用byte数组，减少内存消耗
// 解法二
func countBefore(s []byte) []byte {
	resStr := []byte{}
	// 本质上是对每个数字用重复次数+重复的数字然后拼起来

	// 49就是ascii码中的1
	var repeatCount byte = 49
	for i := 0; i < len(s); i++ {
		if i+1 < len(s) && s[i] == s[i+1] {
			// 寻找出现重复的数字，累计重复次数
			// 如果运行到末尾，且末尾不是重复数字的情况也是执行else的逻辑
			// 如果末尾和之前的重复，那么会被repeatCount记录
			repeatCount++
		} else {

			resStr = append(resStr, repeatCount, s[i])
			// 重置重复次数
			repeatCount = 49
		}
	}
	return resStr
}
func countAndSay(n int) string {
	// temp := []byte("1")
	temp := []byte{49}
	for i := 1; i < n; i++ {
		temp = countBefore(temp)
	}
	return string(temp)
}

func main() {

	// fmt.Println(countBefore("1"))
	// fmt.Println(countBefore("11"))
	// fmt.Println(countBefore("21"))
	// fmt.Println(countBefore("1211"))
	fmt.Println(countAndSay(2))
	fmt.Println(countAndSay(3))
	fmt.Println(countAndSay(4))
}
