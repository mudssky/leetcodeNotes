package main

import "fmt"

/*
Problem:
Given n non-negative integers a1, a2, ..., an , where each represents a point at coordinate (i, ai). n vertical lines are drawn such that the two endpoints of line i is at (i, ai) and (i, 0). Find two lines, which together with x-axis forms a container, such that the container contains the most water.

Note: You may not slant the container and n is at least 2.

The above vertical lines are represented by array [1,8,6,2,5,4,8,3,7]. In this case, the max area of water (blue section) the container can contain is 49.


Example:

Input: [1,8,6,2,5,4,8,3,7]
Output: 49
翻译：
给你n个非负整数 a1,a2,...an,每个数代表坐标中的一个点(i,ai)。在坐标内画n条垂直线，
在坐标内画 n 条垂直线，垂直线 i 的两个端点分别为 (i, ai) 和 (i, 0)。找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。
说明：你不能倾斜容器，且 n 的值至少为 2。

*/

// 解法一：暴力求解，
// 容纳水的面积，就是两根垂直于x轴的线其中最短的一条，乘 两条线间的距离得出来的矩形面积
// 算出所有矩形面积，求出其中的最大值
// 时间复杂度O(n^2) 空间复杂度O(1)
/*
Runtime: 488 ms, faster than 29.01% of Go online submissions for Container With Most Water.
Memory Usage: 5.6 MB, less than 12.50% of Go online submissions for Container With Most Water.
*/

/*
func maxArea(height []int) int {
	max := 0
	arrLen := len(height)
	var containerH int
	for i := 0; i < arrLen; i++ {
		for j := i + 1; j < arrLen; j++ {
			if height[i] > height[j] {
				containerH = height[j]
			} else {
				containerH = height[i]
			}
			containerArea := containerH * (j - i)
			if containerArea > max {
				max = containerArea
			}
		}
	}
	return max
}*/
// 解法二：双指针法，
// 解题的关键之处在于，两线段之间形成的区域总是会收到其中较短那条线段长度的限制，两线段距离越远得到的面积就越大
/*
最初我们考虑由最外围两条线段构成的区域。现在，为了使面积最大化，我们需要考虑更长的两条线段之间的区域。
如果我们试图将指向较长线段的指针向内侧移动，矩形区域的面积将受限于较短的线段而不会获得任何增加。
但是，在同样的条件下，移动指向较短线段的指针尽管造成了矩形宽度的减小，但却可能会有助于面积的增大。
因为移动较短线段的指针会得到一条相对较长的线段，这可以克服由宽度减小而引起的面积减小。

也就是说从两边开始，两边的指针始终是往面积可能更大的方向移动，所以说最大值会在这之中。
此时相遇后再继续移动也只是对称的操作，不会有其他可能性
*/

// 时间复杂度O(n) 空间复杂度O(1)
func maxArea(height []int) int {
	var max int
	left := 0
	right := len(height) - 1
	for left < right {
		if height[left] < height[right] {
			newArea := height[left] * (right - left)
			if newArea > max {
				max = newArea
			}
			left++
		} else {
			newArea := height[right] * (right - left)
			if newArea > max {
				max = newArea
			}
			right--
		}

	}
	return max
}

func main() {
	input := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	fmt.Println(maxArea(input))
}
