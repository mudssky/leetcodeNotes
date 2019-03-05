package main

/*
Problem:
Determine whether an integer is a palindrome. An integer is a palindrome when it reads the same backward as forward.

Example 1:

Input: 121
Output: true
Example 2:

Input: -121
Output: false
Explanation: From left to right, it reads -121. From right to left, it becomes 121-. Therefore it is not a palindrome.
Example 3:

Input: 10
Output: false
Explanation: Reads 01 from right to left. Therefore it is not a palindrome.
Follow up:

Coud you solve it without converting the integer to a string?

翻译：
判断一个整数是否是回文数，如果一个数正着读或是反过来读都是一样的，就是回文数。

提升：
思考一下，能否不把整数转换为字符串，解决这个问题
*/
// 解法一：因为之前判断最大回文子串的时候，写过一个判断回文字符串的函数，直接拿过来用，添加一部把整数转换为字符串
// 时间复杂度：O(n),空间复杂度 O(log n) 因为保存了一个字符串
/*
Runtime: 44 ms, faster than 100.00% of Go online submissions for Palindrome Number.
Memory Usage: 5.2 MB, less than 49.56% of Go online submissions for Palindrome Number.
*/
/*func isPalindrome(x int) bool {
	s := strconv.Itoa(x)
	strLen := len(s)
	for i := 0; i < strLen/2; i++ {
		if s[i] != s[strLen-i-1] {
			return false
		}
	}
	return true
}*/

// 解法二 因为之前写过一个颠倒整数的函数，同样直接拿来用
// 效率比法一差不多，因为法一保存了一个字符串，会占用更多空间
/*
Runtime: 44 ms, faster than 100.00% of Go online submissions for Palindrome Number.
Memory Usage: 5 MB, less than 98.23% of Go online submissions for Palindrome Number.
*/
/*func reverse(x int) int {
	resNum := 0
	for x != 0 {
		last := x % 10
		x /= 10
		resNum = resNum*10 + last
	}
	return resNum
}
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	if reverse(x) == x {
		return true
	}
	return false

}
*/

// 解法三：解法二翻转了整个数字，实际上只要翻转一半进行判断即可，算是对解法二进行了优化
func isPalindrome(x int) bool {
	// 先排除特殊情况
	// 负数不可能是回文数，末尾为0且不是0本身的也进行排除
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}
	var revertedNum int
	// 如果是回文数，进行到一个位数，就会出现两边的数相等或者减少一位相等
	// 这时退出，对这两个数进行比较
	for x > revertedNum {
		last := x % 10
		x /= 10
		revertedNum = revertedNum*10 + last
	}

	return x == revertedNum || x == revertedNum/10

}

func main() {

}
