package main

import (
	"fmt"
)

/*
Implement next permutation, which rearranges numbers into the lexicographically next greater permutation of numbers.

If such arrangement is not possible, it must rearrange it as the lowest possible order (ie, sorted in ascending order).

The replacement must be in-place and use only constant extra memory.

Here are some examples. Inputs are in the left-hand column and its corresponding outputs are in the right-hand column.

1,2,3 → 1,3,2
3,2,1 → 1,2,3
1,1,5 → 1,5,1

实现获取下一个排列的函数，算法需要将给定数字序列重新排列成字典序中下一个更大的排列。

如果不存在下一个更大的排列，则将数字重新排列成最小的排列（即升序排列）。

必须原地修改，只允许使用额外常数空间。

以下是一些例子，输入位于左侧列，其相应输出位于右侧列。
1,2,3 → 1,3,2
3,2,1 → 1,2,3
1,1,5 → 1,5,1

*/
// 解法一
// 所谓字典序这个次，以前做字符串相关的题经常碰到，数字的情况下，字典序和数字的大小顺序一致。所以也可以看成所有数字排列的大小顺序
// 这里就相当于比较这两个字符串拼接起来的数字大小
/*首先，我们观察到对于任何给定序列的降序，没有可能的下一个更大的排列。
例如，以下数组不可能有下一个排列：
[9, 5, 4, 3, 1]
我们需要从右边找到第一对两个连续的数字 a[i] 和 a[i-1]，它们满足 a[i]>a[i-1]。现在，没有对 a[i-1]右侧的重新排列可以创建更大的排列，因为该子数组由数字按降序组成。
因此，我们需要重新排列 a[i-1]a[i−1] 右边的数字，包括它自己。
现在，什么样的重新排列将产生下一个更大的数字？我们想要创建比当前更大的排列。因此，我们需要将数字 a[i-1]a[i−1] 替换为位于其右侧区域的数字中比它更大的数字，例如 a[j]。
我们交换数字 a[i-1] 和a[j]。我们现在在索引i−1 处有正确的数字。 但目前的排列仍然不是我们正在寻找的排列。我们需要通过仅使用a[i-1]右边的数字来形成最小的排列。 因此，我们需要放置那些按升序排列的数字，以获得最小的排列。

但是，请记住，在从右侧扫描数字时，我们只是继续递减索引直到我们找到 a[i]和a[i-1]这对数。其中，a[i] > a[i-1]因此，
a[i−1] 右边的所有数字都已按降序排序。此外，交换 a[i-1]和 a[j]并未改变该顺序。因此，我们只需要反转 a[i-1] 之后的数字，以获得下一个最小的字典排列。
*/

// 一遍扫描，从右边往左边扫描，如果，存在降序的序列，此时把这个小的数字，交换右侧更大的数字中最小的。
// 因为右边相当于个位，要想得到稍微大一点的数，肯定是要动低位数，如果在高位数中找更大的交换，高位数，就会变小，那么整体会不可避免地变小
// 之后我们要把整个右边从小到大排序，因为交换a[i-1]和右边最小的，已经得到了比原来更大的一个数，但不是最小，只有右边是从小到大排列时才是最小
/*func nextPermutation(nums []int) {
	numsLen := len(nums)
	min := -1
	for i := numsLen - 1; i >= 1; i-- {
		if nums[i] > nums[i-1] {
			min = i
			j := i
			for ; j < numsLen; j++ {
				if nums[j] > nums[i-1] && nums[j] < nums[min] {
					min = j
				}
			}
			nums[i-1], nums[min] = nums[min], nums[i-1]
			sort.Ints(nums[i:])
			break
		}
	}
	if min == -1 {
		sort.Ints(nums)
	}
}
*/
// 解法一优化
// 调换两个元素位置后，后面部分的排序，其实只要把那部分数组位置调过来就可以，因为经过前面的扫描，后面已经是降序了。
// 我们只要调过来就是升序，就是最小的
// 同理,后面min=-1时的排序,也是将数组掉头的一个操作
func reverse(nums []int) {
	numsLen := len(nums)
	for i, j := 0, numsLen-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
}
func nextPermutation(nums []int) {
	numsLen := len(nums)
	min := -1
	for i := numsLen - 1; i >= 1; i-- {
		if nums[i] > nums[i-1] {
			min = i
			j := i
			for ; j < numsLen; j++ {
				if nums[j] > nums[i-1] && nums[j] <= nums[min] {
					min = j
				}
			}
			nums[i-1], nums[min] = nums[min], nums[i-1]
			reverse(nums[i:])
			break
		}
	}
	if min == -1 {
		reverse(nums[:])
	}
}
func main() {
	// input := []int{1, 2, 3}
	input := []int{2, 3, 1, 3, 3}
	// fmt.Println(input)
	nextPermutation(input)
	fmt.Println(input)
}
