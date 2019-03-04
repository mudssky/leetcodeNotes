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