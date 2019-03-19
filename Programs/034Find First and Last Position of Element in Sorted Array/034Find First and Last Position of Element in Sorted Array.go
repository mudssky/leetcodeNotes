package main

import "fmt"

/*

Given an array of integers nums sorted in ascending order, find the starting and ending position of a given target value.

Your algorithm's runtime complexity must be in the order of O(log n).

If the target is not found in the array, return [-1, -1].

Example 1:

Input: nums = [5,7,7,8,8,10], target = 8
Output: [3,4]
Example 2:

Input: nums = [5,7,7,8,8,10], target = 6
Output: [-1,-1]

给定一个按照升序排列的整数数组 nums，和一个目标值 target。找出给定目标值在数组中的开始位置和结束位置。

你的算法时间复杂度必须是 O(log n) 级别。

如果数组中不存在目标值，返回 [-1, -1]。

示例 1:

输入: nums = [5,7,7,8,8,10], target = 8
输出: [3,4]
示例 2:

输入: nums = [5,7,7,8,8,10], target = 6
输出: [-1,-1]

*/
// 解法一：很明显又是一道二分查找题
// 执行两次二分查找即可，只是要稍微做点变动
// 第一次查找找到最左边的数,第二次查找找最右边的数
// 举例说明,例如找第一个相等的数,二分查找如果找到相等的数,并不马上返回,而是判断左边还有没有数,如果没有数,说明这个相等的数下标已经是
// 最左边了。如果有数，判断是否相等，如果相等，则对mid-1的区域继续进行二分查找。不断向左缩小范围，直到只能找到一个数
// 找右边的数也是同理
func findFirst(nums []int, low, high, target int) int {
	// 退出条件 low>high说明数组已经找遍了
	if low > high {
		return -1
	}
	mid := (low + high) / 2
	// 如果mid下标的值和目标target相等，我们还要判断左边还有没有相等的
	if nums[mid] == target {
		// 如果已经是最左边，那么直接返回0
		if mid == 0 {
			return 0
		}
		// 不是最左边，并且左边的mid-1存在相等，缩小查找范围
		if nums[mid-1] == target {
			return findFirst(nums, low, mid-1, target)
		}
		// 如果mid-1不相等，说明已经是最左边，直接返回
		return mid
	}

	if target < nums[mid] && target >= nums[low] {
		return findFirst(nums, low, mid-1, target)
	} else if target > nums[mid] && target <= nums[high] {
		return findFirst(nums, mid+1, high, target)
	} else {
		return -1
	}
}

func findLast(nums []int, low, high, target int) int {
	if low > high {
		return -1
	}
	mid := (low + high) / 2
	if nums[mid] == target {
		if mid == len(nums)-1 {
			return high
		}
		if nums[mid+1] == target {
			return findLast(nums, mid+1, high, target)
		}
		return mid
	}

	if target < nums[mid] && target >= nums[low] {
		return findLast(nums, low, mid-1, target)
	} else if target > nums[mid] && target <= nums[high] {
		return findLast(nums, mid+1, high, target)
	} else {
		return -1
	}
}
func searchRange(nums []int, target int) []int {
	p1 := findFirst(nums, 0, len(nums)-1, target)
	p2 := findLast(nums, 0, len(nums)-1, target)
	return []int{p1, p2}
}

func main() {
	// input := []int{5, 7, 7, 8, 8, 10}
	input := []int{2, 2}
	// fmt.Println(searchRange(input, 8))
	fmt.Println(searchRange(input, 2))
}
