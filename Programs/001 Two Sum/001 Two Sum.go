package main

import "fmt"

/*
Given an array of integers, return indices of the two numbers such that they add up to a specific target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

Example:

Given nums = [2, 7, 11, 15], target = 9,

Because nums[0] + nums[1] = 2 + 7 = 9,
return [0, 1].

*/

/*
翻译
给你一个整数的数组，返回相加能得到目标数字的元素的索引
假定每个输入只有一个解决方案，而且同一个元素不会用到两次
*/

/**
 * @param {number[]} nums
 * @param {number} target
 * @return {number[]}
 */
//  js解法
//  var twoSum = function(nums, target) {
// 	for( let i=0;i<nums.length;i++){
// 		for(let j=i+1;j<nums.length;j++){
// 			if(nums[i]+nums[j]==target){
// 				let res = [];
// 				res.push(i);
// 				res.push(j);
// 				return res;
// 			}
// 		}
// 	}
//    return;
// };

// 解法一：暴力求解
// 时间复杂度O(n^2) 空间复杂度O(1)
/*
func twoSum(nums []int, target int) []int {
	res := make([]int, 2)
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				res[0] = i
				res[1] = j
				return res
			}
		}
	}
	return nil
}
*/

// 解法二，用哈希表存储数据，遍历两次，第一次把数据放入哈希表中，第二次遍历查找哈希表
// 由于哈希表查找近似于查找数组，查找时间复杂度接近O(1)
// 两次遍历为长度n的数组，时间复杂度为O(n),空间复杂度为O(n)，因为创建了长度为n的哈希表和数组
// 因为go自带的map不是按顺序遍历的,所以用一个数组来记录每个键的顺序
/*func twoSum(nums []int, target int) []int {
	dict := make(map[int]int)
	var keySorted []int
	for i := 0; i < len(nums); i++ {
		dict[nums[i]] = i
		keySorted = append(keySorted, nums[i])
	}
	//go里面range遍历map是随机的，所以要自己写遍历
	for i, v := range keySorted {
		fmt.Println(i, v)
		value, ok := dict[target-v]
		if ok && value != i {
			return []int{i, value}
		}
	}

	// for i, v := range dict {
	// 	value, ok := dict[target-i]
	// 	if ok {
	// 		return []int{v, value}
	// 	}
	// }
	return nil
}*/
// 犯了个错误，原来给的数组是可以有重复的啊,不过上面的方法因为保存了一个keySorted数组
// ,倒着用map的时候即使即使覆盖了前面的数据,因为覆盖数字永远比较大,顺序不会出错,所以还是通过了测试

// 解法三：解法二中用了两次map遍历，并且多用了一个数组存储哈希表的顺序。
// 其实查找的过程可以在添加哈希表的过程中完成
// 时间复杂度O(n),空间复杂度O(n)
/*
可见国外没几个用golang做题的
Runtime: 4 ms, faster than 100.00% of Go online submissions for Two Sum.
Memory Usage: 3.7 MB, less than 51.62% of Go online submissions for Two Sum.

国内是这样
执行用时: 4 ms, 在Two Sum的Go提交中击败了100.00% 的用户
内存消耗: 3.7 MB, 在Two Sum的Go提交中击败了2.35% 的用户


*/

func twoSum(nums []int, target int) []int {
	dict := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		if v, ok := dict[target-nums[i]]; ok {
			return []int{v, i}
		}
		dict[nums[i]] = i
	}
	return nil
}
func main() {
	fmt.Println(twoSum([]int{3, 2, 4}, 6))
}
