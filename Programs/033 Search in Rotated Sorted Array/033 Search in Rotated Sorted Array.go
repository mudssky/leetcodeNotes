package main

import "fmt"

/*
Suppose an array sorted in ascending order is rotated at some pivot unknown to you beforehand.

(i.e., [0,1,2,4,5,6,7] might become [4,5,6,7,0,1,2]).

You are given a target value to search. If found in the array return its index, otherwise return -1.

You may assume no duplicate exists in the array.

Your algorithm's runtime complexity must be in the order of O(log n).

Example 1:

Input: nums = [4,5,6,7,0,1,2], target = 0
Output: 4
Example 2:

Input: nums = [4,5,6,7,0,1,2], target = 3
Output: -1

假设按照升序排序的数组在预先未知的某个点上进行了旋转。

( 例如，数组 [0,1,2,4,5,6,7] 可能变为 [4,5,6,7,0,1,2] )。

搜索一个给定的目标值，如果数组中存在这个目标值，则返回它的索引，否则返回 -1 。

你可以假设数组中不存在重复的元素。

你的算法时间复杂度必须是 O(log n) 级别。

示例 1:

输入: nums = [4,5,6,7,0,1,2], target = 0
输出: 4
示例 2:

输入: nums = [4,5,6,7,0,1,2], target = 3
输出: -1
*/
// 解法一：
// 这题的问题在于，要求时间复杂度O(log n),很明显是让我们用二分查找的思路做题
// 评论区一位老哥说的很好很清晰
// 将数组一分为二，其中一定有一个是有序的，另一个可能是有序，也能是部分有序。此时有序部分用二分法查找。无序部分再一分为二，其中一个一定有序，另一个可能有序，可能无序。就这样循环.
// 那么怎么判断数组有序还是无序呢？很明显，比较开头的值和中间的值即可。若有序，则开头值小于中间值
/*
func binSearch(nums []int, target int) int {
	low := 0
	high := len(nums) - 1

	for low <= high {
		mid := (low + high) / 2
		if nums[mid] == target {
			return mid
		} else if nums[mid] < target {
			low = mid + 1
		} else if nums[mid] > target {
			high = mid - 1
		}
	}
	return -1
}

func search(nums []int, target int) int {
	numsLen := len(nums)
	if numsLen == 0 {
		return -1
	}
	mid := (numsLen - 1) / 2
	if nums[mid] == target {
		return mid
	}
	if nums[0] <= nums[mid/2] {
		if inx := search(nums[mid+1:], target); inx != -1 {
			return inx
		}
		if index := binSearch(nums[:mid], target); index != -1 {
			return index
		}
	} else {
		// 这里用二分查找没把下标传入，相当于传入了一个新数组，所以出错
		if inx := binSearch(nums[mid+1:], target); inx != -1 {
			return inx
		}
		if index := search(nums[:mid], target); index != -1 {
			return index
		}

	}
	return -1
}*/

func binSearch(nums []int, low, high, target int) int {
	if low > high {
		return -1
	}
	mid := (low + high) / 2
	if nums[mid] == target {
		return mid
	}
	// 判断起始点是否小于最高点，小于说明后半部分有序，前半部分无序
	// 这里是二分，只能进入一条逻辑
	if nums[mid] < nums[high] {
		// 如果目标值，在high和mid下标之间, 这时一个有序的区间，保证target在这个区间内进行二分查找
		// 由于二分查找这部分逻辑在上面，所以如果能找到会先返回，if条件如果满足必然能找到。所以就没后面什么事了
		// if条件不满足，就会重新执行二分，判断的过程，即下面一条逻辑
		if nums[mid] < target && target <= nums[high] {
			return binSearch(nums, mid+1, high, target)
		}

		return binSearch(nums, low, mid-1, target)
	} else {
		// 此时，low-mid的区间是有序的
		// 如果目标值，在low和mid下标之间, 这时一个有序的区间，保证target在这个区间内进行二分查找
		if nums[low] <= target && target < nums[mid] {
			return binSearch(nums, low, mid-1, target)
		}
		return binSearch(nums, mid+1, high, target)
	}

}

func search(nums []int, target int) int {
	return binSearch(nums, 0, len(nums)-1, target)
}
func main() {
	input := []int{1, 3}
	fmt.Println(search(input, 2))
}
