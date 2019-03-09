package main

/*
Given an array nums of n integers and an integer target, are there elements a, b, c, and d in nums such that a + b + c + d = target? Find all unique quadruplets in the array which gives the sum of target.

Note:

The solution set must not contain duplicate quadruplets.

Example:

Given array nums = [1, 0, -1, 0, -2, 2], and target = 0.

A solution set is:
[
  [-1,  0, 0, 1],
  [-2, -1, 1, 2],
  [-2,  0, 0, 2]
]

给定一个包含 n 个整数的数组 nums 和一个目标值 target，判断 nums 中是否存在四个元素 a，b，c 和 d ，使得 a + b + c + d 的值与 target 相等？找出所有满足条件且不重复的四元组。

注意：

答案中不可以包含重复的四元组。

示例：

给定数组 nums = [1, 0, -1, 0, -2, 2]，和 target = 0。

满足要求的四元组集合为：
[
  [-1,  0, 0, 1],
  [-2, -1, 1, 2],
  [-2,  0, 0, 2]
]

*/
// 解法一，很明显在3Sum的基础上，再加一层循环就可以了

import "sort"

func fourSum(nums []int, target int) [][]int {
	ans := make([][]int, 0)
	ell := len(nums)

	sort.Ints(nums)
	for i := 0; i < ell-3; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		for j := i + 1; j < ell-2; j++ {
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}

			for k, l := j+1, ell-1; k < l; {
				sum := nums[i] + nums[j] + nums[k] + nums[l]
				if sum == target {
					ans = append(ans, []int{nums[i], nums[j], nums[k], nums[l]})
					k++
					l--
					for k < l && nums[k] == nums[k-1] {
						k++
					}
					for k < l && nums[l] == nums[l+1] {
						l--
					}
				} else if sum < target {
					k++
				} else {
					l--
				}
			}
		}
	}

	return ans
}
func main() {

}
