package main

import "fmt"

/*
Given a sorted array and a target value, return the index if the target is found. If not, return the index where it would be if it were inserted in order.

You may assume no duplicates in the array.

Example 1:

Input: [1,3,5,6], 5
Output: 2
Example 2:

Input: [1,3,5,6], 2
Output: 1
Example 3:

Input: [1,3,5,6], 7
Output: 4
Example 4:

Input: [1,3,5,6], 0
Output: 0

给定一个排序数组和一个目标值，在数组中找到目标值，并返回其索引。如果目标值不存在于数组中，返回它将会被按顺序插入的位置。

你可以假设数组中无重复元素。

示例 1:

输入: [1,3,5,6], 5
输出: 2
示例 2:

输入: [1,3,5,6], 2
输出: 1
示例 3:

输入: [1,3,5,6], 7
输出: 4
示例 4:

输入: [1,3,5,6], 0
输出: 0
*/

// 解法一：这题是简单题，对时间复杂度没有要求，所以用普通的遍历也可以
// 时间复杂度O(n) 空间复杂度O(1)
/*func searchInsert(nums []int, target int) int {
	numsLen := len(nums)
	if numsLen == 0 || nums[0] > target {
		return 0
	}
	if nums[0] == target {
		return 0
	}
	for i := 0; i < numsLen-1; i++ {
		if nums[i] == target {
			return i
		} else if nums[i] < target && target <= nums[i+1] {
			return i + 1
		}
	}
	return numsLen
}
*/
// 解法一简化一下
/*
func searchInsert(nums []int, target int) int {
	numsLen := len(nums)
	for i := 0; i < numsLen; i++ {
		if nums[i] >= target {
			return i
		}
	}
	return numsLen
}
*/
// 解法二 运用二分查找，
// 先把查找的范围简化一下，排除不在数组范围的情况
/*Runtime: 4 ms, faster than 100.00% of Go online submissions for Search Insert Position.
Memory Usage: 3 MB, less than 100.00% of Go online submissions for Search Insert Position.*/
func binSearch(nums []int, low, high, target int) int {
	if low > high {
		return -1
	}
	mid := (low + high) / 2
	if nums[mid] == target {
		return mid
	} else if target > nums[mid] && target < nums[mid+1] {
		return mid + 1
	} else if target > nums[mid]-1 && target < nums[mid] {
		return mid
	}

	if target < nums[mid] && target >= nums[low] {
		return binSearch(nums, low, mid-1, target)
	}

	return binSearch(nums, mid+1, high, target)

}

func searchInsert(nums []int, target int) int {
	numsLen := len(nums)
	if numsLen == 0 || nums[0] > target {
		return 0
	}
	if nums[numsLen-1] < target {
		return numsLen
	}
	return binSearch(nums, 0, numsLen-1, target)
}
func main() {
	// input := []int{1, 3, 5, 6}
	input := []int{1}
	fmt.Println(searchInsert(input, 1))
}
