package main

import (
	"fmt"
	"sort"
)

/*
Problem:
Given an array nums of n integers, are there elements a, b, c in nums such that a + b + c = 0? Find all unique triplets in the array which gives the sum of zero.

Note:

The solution set must not contain duplicate triplets.

Example:

Given array nums = [-1, 0, 1, 2, -1, -4],

A solution set is:
[
  [-1, 0, 1],
  [-1, -1, 2]
]

翻译：
给你一个n个整数的数组，求是否存在a，b，c在这些整数中使得a+b+c=0，找到所有不重复的3元组
*/
// 解法一：
// 先找到3个元素的所有不同组合，然后判断这些组合是否等于0，等于0的加入到结果数组中，时间复杂度O(n^3) 空间复杂度O(n^2)

// 如果返回两个元素的所有排列组合，写法是下面这样，3个数只需，再加一层循环
func twoArranges(arr []int) [][]int {
	res := [][]int{}
	arrLen := len(arr)
	for i := 0; i < arrLen-1; i++ {
		for j := i + 1; j < arrLen; j++ {
			res = append(res, []int{arr[i], arr[j]})
		}
	}
	return res
}

func threeArranges(arr []int) [][]int {
	res := [][]int{}
	arrLen := len(arr)
	for k1 := 0; k1 < arrLen-2; k1++ {
		for k2 := k1 + 1; k2 < arrLen-1; k2++ {
			for k3 := k2 + 1; k3 < arrLen; k3++ {
				res = append(res, []int{arr[k1], arr[k2], arr[k3]})
			}
		}
	}

	return res
}

/*
type arrLen3 struct {
	num1, num2, num3 int
}

// 用一个map结构列出3个数所有的排列方式，用于查表
func arrDictLen3(arr []int) map[arrLen3]bool {
	arrDict := make(map[arrLen3]bool)
	arrDict[arrLen3{arr[0], arr[1], arr[2]}] = true
	arrDict[arrLen3{arr[0], arr[2], arr[1]}] = true
	arrDict[arrLen3{arr[1], arr[0], arr[2]}] = true
	arrDict[arrLen3{arr[1], arr[2], arr[0]}] = true
	arrDict[arrLen3{arr[2], arr[0], arr[1]}] = true
	arrDict[arrLen3{arr[2], arr[1], arr[0]}] = true
	return arrDict
}

func threeSum(nums []int) [][]int {
	res := [][]int{}
	numsLen := len(nums)
	for k1 := 0; k1 < numsLen-2; k1++ {
		for k2 := k1 + 1; k2 < numsLen-1; k2++ {
			for k3 := k2 + 1; k3 < numsLen; k3++ {
				if nums[k1]+nums[k2]+nums[k3] == 0 {
					// 判断和为0之后，还要基于res数组判断是否重复
					repeated := false
					// 获取3个数的所有排列
					arrDict := arrDictLen3([]int{nums[k1], nums[k2], nums[k3]})
					// 遍历已经产生的结果，如果没有和这3个数的所有排列一样的，说明没有重复
					for i := 0; i < len(res); i++ {

						if _, ok := arrDict[arrLen3{res[i][0], res[i][1], res[i][2]}]; ok {
							repeated = true
							break
						}
					}
					if repeated == true {
						continue
					}

					res = append(res, []int{nums[k1], nums[k2], nums[k3]})
				}
			}
		}
	}
	return res
}*/
// 解法二：只需要遍历两个数，第三个数用map来查，查不到就不存在
/*
type arrLen3 struct {
	num1, num2, num3 int
}

// 用一个map结构列出3个数所有的排列方式，用于查表
// 判断重复的循环也不需要了，直接用结果6倍大小的哈希表查询
func threeSum(nums []int) [][]int {
	// 用于存储结果
	res := [][]int{}
	// 字典，用于查询第三个数，节省一层循环
	dict := make(map[int]int)
	for i, v := range nums {
		dict[v] = i
	}
	numsLen := len(nums)
	// 排列字典，用于保存3个数的所有排列，每得到一个新答案，就把答案的6个结果放进去，用于查表
	numsDict := make(map[arrLen3]bool)
	for k1 := 0; k1 < numsLen-1; k1++ {
		for k2 := k1 + 1; k2 < numsLen; k2++ {
			k3Val := 0 - nums[k1] - nums[k2]
			// 除了能从表中查出以外，也就是满足a+b+c==0，k3也就是下标还不能和k1，k2一样
			if k3, ok := dict[k3Val]; ok && k3 != k1 && k3 != k2 {
				if _, ook := numsDict[arrLen3{nums[k1], nums[k2], nums[k3]}]; ook {
					continue
				} else {
					numsDict[arrLen3{nums[k1], nums[k2], nums[k3]}] = true
					numsDict[arrLen3{nums[k1], nums[k3], nums[k2]}] = true
					numsDict[arrLen3{nums[k2], nums[k1], nums[k3]}] = true
					numsDict[arrLen3{nums[k2], nums[k3], nums[k1]}] = true
					numsDict[arrLen3{nums[k3], nums[k1], nums[k2]}] = true
					numsDict[arrLen3{nums[k3], nums[k2], nums[k1]}] = true
					res = append(res, []int{nums[k1], nums[k2], nums[k3]})
				}
			}
		}
	}
	return res
}*/

// 解法三，优化解法二，如果数组排好序，有可能一次遍历搞定吗？
// 用双指针法，时间复杂度仍然接近O(n^2),但是空间复杂度降低不少。
// 相当于把数组排序分成正数负数两个部分，nums[i]随着i增大逐渐往正数区靠近
// 但是到正数区之后就不用判断了，因为此时3个数全部大于0了。时间复杂度其实是O(+*-)
func threeSum(nums []int) [][]int {
	if len(nums) < 3 {
		return nil
	}
	res := make([][]int, 0)
	sort.Ints(nums)
	for i := 0; i < len(nums) && nums[i] <= 0; i++ {
		// i指针指向的数字如果已经出现过一次，之后的都是重复的情况
		// 3数和i移动时候最多容忍两个数重复，因为如果两个数重复第三个数就确定了
		// 因为指针移动的缘故，判断到两个数重复的时候，之前 i=? l=?的情况已经计算过了。
		// 为了达到去重的效果所以把重复的部分都跳过，最多重复两个（第三个不判断，但是也可能重复），且过滤到只出现一次这种情况
		// 因为两个数确定，第三个数也就跟着确定，所以第三个数不用判断，两个重复就是等于3个形成的序列重复了，
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		l, r := i+1, len(nums)-1
		for l < r {
			if nums[l]+nums[r] == (-nums[i]) {

				if r == len(nums)-1 || nums[r+1] != nums[r] {
					res = append(res, []int{nums[i], nums[l], nums[r]})
				}

				l++
				r--
			}
			if nums[l]+nums[r] < (-nums[i]) {
				l++
			}
			if nums[l]+nums[r] > (-nums[i]) {
				r--
			}
		}

	}

	return res
}
func main() {
	// nums := []int{-1, 0, 1, 2, -1, -4}
	nums := []int{-1, -2, -3, 4, 1, 3, 0, 3, -2, 1, -2, 2, -1, 1, -5, 4, -3}
	fmt.Println(twoArranges(nums))
	// fmt.Println(threeArranges(nums))
	fmt.Println(threeSum(nums))
}
