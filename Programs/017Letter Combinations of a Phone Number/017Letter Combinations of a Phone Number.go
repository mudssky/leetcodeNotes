package main

import "fmt"

/*
Given a string containing digits from 2-9 inclusive, return all possible letter combinations that the number could represent.

A mapping of digit to letters (just like on the telephone buttons) is given below. Note that 1 does not map to any letters.



Example:

Input: "23"
Output: ["ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"].
Note:

Although the above answer is in lexicographical order, your answer could be in any order you want.


给定一个仅包含数字 2-9 的字符串，返回所有它能表示的字母组合。

给出数字到字母的映射如下（与电话按键相同）。注意 1 不对应任何字母。

示例:

输入："23"
输出：["ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"].
说明:
尽管上面的答案是按字典序排列的，但是你可以任意选择答案输出的顺序。
*/

// 解法一，回朔法
/*
func find(digits string, index int, s string, dict map[byte]string, res *[]string) {
	if index == len(digits) {
		*res = append(*res, s)
	} else {
		numStr := digits[index]
		for i := 0; i < len(dict[numStr]); i++ {
			find(digits, index+1, s+string(dict[numStr][i]), dict, res)
		}
	}
}

func letterCombinations(digits string) []string {
	res := make([]string, 0)

	if len(digits) == 0 {
		return res
	}

	var dict = map[byte]string{
		'2': "abc",
		'3': "def",
		'4': "ghi",
		'5': "jkl",
		'6': "mno",
		'7': "pqrs",
		'8': "tuv",
		'9': "wxyz",
	}
	find(digits, 0, "", dict, &res)
	return res

}*/

// 解法二：把递归改成循环

func letterCombinations(digits string) []string {
	res := make([]string, 0)
	digitWordsMap := map[string]string{
		"2": "abc",
		"3": "def",
		"4": "ghi",
		"5": "jkl",
		"6": "mno",
		"7": "pqrs",
		"8": "tuv",
		"9": "wxyz",
	}
	for _, digit := range digits {
		// 查表找到数字对应的字母
		words := digitWordsMap[string(digit)]
		tmp := make([]string, 0)
		// 遍历数字对应的字母，结果集中还没有字符串，就把，遍历到的字母字符分别放进去
		for _, word := range words {
			if len(res) > 0 {
				for _, item := range res {
					// 遍历上次的结果，把这次的字母拼接上，用tmp暂存
					tmp = append(tmp, item+string(word))
				}
			} else {
				// 第一次的情况下，直接把字母作为字符串插入
				tmp = append(tmp, string(word))
			}
		}
		res = tmp
	}
	return res
}

func main() {
	digits := "23"
	fmt.Println(letterCombinations(digits))

}
