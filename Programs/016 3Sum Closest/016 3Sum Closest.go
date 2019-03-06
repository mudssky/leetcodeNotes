package main

import (
	"fmt"
	"sort"
)

/*
Given an array nums of n integers and an integer target, find three integers in nums such that the sum is closest to target. Return the sum of the three integers. You may assume that each input would have exactly one solution.

Example:

Given array nums = [-1, 2, 1, -4], and target = 1.

The sum that is closest to the target is 2. (-1 + 2 + 1 = 2).

给定一个包括 n 个整数的数组 nums 和 一个目标值 target。找出 nums 中的三个整数，使得它们的和与 target 最接近。返回这三个数的和。假定每组输入只存在唯一答案。

例如，给定数组 nums = [-1，2，1，-4], 和 target = 1.

与 target 最接近的三个数的和为 2. (-1 + 2 + 1 = 2).
*/
// 解法一：和上一道题一样，此题显然也可以用双指针法。
// 但是因为排序过于耗时了，很可能不是最优的解法
// 时间复杂度O(n^2)
func getAbs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}
func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	fmt.Println(nums)
	numsLen := len(nums)
	closestNum := nums[0] + nums[1] + nums[2]
	if closestNum >= target {
		return closestNum
	}
	for i := 0; i < numsLen-2; i++ {
		l := i + 1
		r := numsLen - 1
		for l < r {
			threeSum := nums[l] + nums[r] + nums[i]
			if getAbs(threeSum-target) < getAbs(closestNum-target) {
				closestNum = threeSum
			}
			if threeSum == target {
				return target
			} else if threeSum < target {
				l++
			} else {
				// 如果已经等于目标的话，直接返回
				r--
			}
		}
	}
	return closestNum
}

func main() {
	inputNums := []int{-1, 2, 1, -4}
	fmt.Println(threeSumClosest(inputNums, 1))
}
