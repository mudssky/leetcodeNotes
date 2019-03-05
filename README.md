# leetcodeNotes
use golang solve problems in leetcode



## 001 Two Sum

Given an array of integers, return indices of the two numbers such that they add up to a specific target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

Example:

Given nums = [2, 7, 11, 15], target = 9,

Because nums[0] + nums[1] = 2 + 7 = 9,

return [0, 1].

翻译

给你一个整数的数组，返回相加能得到目标数字的元素的索引

假定每个输入只有一个解决方案，而且同一个元素不会用到两次

## 002Add Two Numbers

Problem:

You are given two non-empty linked lists representing two non-negative integers. The digits are stored in reverse order and each of their nodes contain a single digit. Add the two numbers and return it as a linked list.

You may assume the two numbers do not contain any leading zero, except the number 0 itself.

Example:

Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)

Output: 7 -> 0 -> 8

Explanation: 342 + 465 = 807.

翻译：

给你两个非空链表代表两个非负整数，每一位数字都都是按照逆序存储的，把这两个数相加，并且用链表形式返回

假定两个数开头都不包含0，除了他们自己是0的时候

## 003Longest Substring Without Repeating Characters

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

​             Note that the answer must be a substring, "pwke" is a subsequence and not a substring.

翻译：

给你一个字符串，找到最长子串求出长度，不能有重复的字母



## 004Median of Two Sorted Arrays

==待续==

==TODO==

There are two sorted arrays nums1 and nums2 of size m and n respectively.

Find the median of the two sorted arrays. The overall run time complexity should be O(log (m+n)).

You may assume nums1 and nums2 cannot be both empty.

Example 1:

nums1 = [1, 3]

nums2 = [2]

The median is 2.0

Example 2:

nums1 = [1, 2]

nums2 = [3, 4]

The median is (2 + 3)/2 = 2.5

翻译：

有两个排序好的数组，大小分别为m和n

给出两个数组的中位数，时间复杂度需要控制在 O(log (m+n))

假定两个数组不能同时为空





## 005Longest Palindromic Substring

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









## 006.ZigZag Conversion

Problem:

The string "PAYPALISHIRING" is written in a zigzag pattern on a given number of rows like this: (you may want to display this pattern in a fixed font for better legibility)

P   A   H   N

A P L S I I G

Y   I   R

And then read line by line: "PAHNAPLSIIGYIR"

Write the code that will take a string and make this conversion given a number of rows:

string convert(string s, int numRows);

Example 1:

Input: s = "PAYPALISHIRING", numRows = 3

Output: "PAHNAPLSIIGYIR"

Example 2:

Input: s = "PAYPALISHIRING", numRows = 4

Output: "PINALSIGYAHRPI"

Explanation:

P     I    N

A   L S  I G

Y A   H R

P     I

翻译：

字符串 "PAYPALISHIRING" 以如下所示的z字形排列

编写代码，输入字符串，和产生字形的行数，打印出对应的z变形

可以看出z的特点是整体占据的区域是一个正方形区域。



## 007Reverse Integer

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





## 008String to Integer (atoi)



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

​             Then take as many numerical digits as possible, which gets 42.

Example 3:

Input: "4193 with words"

Output: 4193

Explanation: Conversion stops at digit '3' as the next character is not a numerical digit.

Example 4:

Input: "words and 987"

Output: 0

Explanation: The first non-whitespace character is 'w', which is not a numerical

​             digit or a +/- sign. Therefore no valid conversion could be performed.

Example 5:

Input: "-91283472332"

Output: -2147483648

Explanation: The number "-91283472332" is out of the range of a 32-bit signed integer.

​             Thefore INT_MIN (−231) is returned.

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



## 009.Palindrome Number

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





## 010.Regular Expression Matching

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

