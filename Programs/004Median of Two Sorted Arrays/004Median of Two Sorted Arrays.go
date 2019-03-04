package main

import (
	"fmt"
	"sort"
)

/*
problem:
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
*/
//解法一:将两个数组排序到一个新数组，这相当于插入排序时间复杂度是O(m+n)（遍历两个数组）,明显达不到题目中要求的
/*func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var i, j, index int
	lenNums1 := len(nums1)
	lenNums2 := len(nums2)
	lenNew := lenNums1 + lenNums2
	newArray := make([]int, lenNew)
	if lenNums1 != 0 && lenNums2 != 0 {
		for i < lenNums1 && j < lenNums2 {
			if nums1[i] < nums2[j] {
				newArray[index] = nums1[i]
				i++
				index++
			} else if nums1[i] > nums2[j] {
				newArray[index] = nums2[j]
				j++
				index++
			}
		}
	}
	if i == lenNums1 && j < lenNums2 {
		for ; j < lenNums2; j++ {
			newArray[index] = nums2[j]
			index++
		}
	} else if j == lenNums2 && i < lenNums1 {
		for ; i < lenNums1; i++ {
			newArray[index] = nums1[i]
			index++
		}
	}
	if lenNew == 1 {
		return float64(newArray[0])
	}
	fmt.Println(newArray)

	if lenNew%2 == 0 {
		return float64(newArray[lenNew/2-1]+newArray[lenNew/2]) / 2.0
	} else {
		return float64(newArray[lenNew/2])
	}
}*/
// 解法二:用go语言自带的排序算法
/*
Runtime: 20 ms, faster than 100.00% of Go online submissions for Median of Two Sorted Arrays.
Memory Usage: 5.8 MB, less than 52.04% of Go online submissions for Median of Two Sorted Arrays.
*/
//时间复杂度，官方的sort如果用的是快排的话，那么时间复杂度nlogn
/*
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	newArray := make([]int, 0)
	newArray = append(newArray, nums1...)
	newArray = append(newArray, nums2...)
	lenNew := len(newArray)
	if lenNew == 1 {
		return float64(newArray[0])
	}
	sort.Slice(newArray, func(i, j int) bool {
		return newArray[i] < newArray[j]
	})
	if lenNew%2 == 0 {
		return float64(newArray[lenNew/2-1]+newArray[lenNew/2]) / 2.0
	} else {
		return float64(newArray[lenNew/2])
	}
}
*/

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	newArray := make([]int, 0)
	newArray = append(newArray, nums1...)
	newArray = append(newArray, nums2...)
	lenNew := len(newArray)
	if lenNew == 1 {
		return float64(newArray[0])
	}
	sort.Slice(newArray, func(i, j int) bool {
		return newArray[i] < newArray[j]
	})
	if lenNew%2 == 0 {
		return float64(newArray[lenNew/2-1]+newArray[lenNew/2]) / 2.0
	} else {
		return float64(newArray[lenNew/2])
	}
}

//官方解法很长一篇，看着头晕，正确的解法还是待续...
func main() {
	// nums1 := []int{1, 3}
	// nums2 := []int{2}
	nums1 := []int{1, 3, 4}
	nums2 := []int{2}
	fmt.Println(findMedianSortedArrays(nums1, nums2))
}
