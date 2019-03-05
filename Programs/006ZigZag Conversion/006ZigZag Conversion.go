package main

import "fmt"

/*
Problem:
The string "PAYPALISHIRING" is written in a zigzag pattern on a given number of rows like this: (you may want to display this pattern in a fixed font for better legibility)

P   A   H   N
A P L S I I G
Y   I   R
And then read line by line: "PAHNAPLSIIGYIR"

Write the code that will take a string and make this conversion given a number of rows:

string convert(string s, int numRows);
Example 1:

Input: s = "PAYPALISHIRING", numRows = 3
Output: "PAHNAPLSIIGYIR"
Example 2:

Input: s = "PAYPALISHIRING", numRows = 4
Output: "PINALSIGYAHRPI"
Explanation:

P     I    N
A   L S  I G
Y A   H R
P     I

翻译：
字符串 "PAYPALISHIRING" 以如下所示的z字形排列
编写代码，输入字符串，和产生字形的行数，打印出对应的z变形

可以看出z的特点是整体占据的区域是一个正方形区域。

*/

// 解法一：找规律把图形循环画出来，每个坐标的字母存到二维数组里，然后遍历数组，拼接字符串
// 用二维数组来记录图形的横坐标纵坐标，以左下角为原点可知，
// 图形的走势是一条y--的直线，以及x++ y++的斜线的组合不断循环。
// 时间复杂度O(n^2)
// 空间复杂度O(n^2)
// 然而问题是会有许多多余的空格，提交的时候一些很长的测试用例会报错 too large file
/*
func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	// 初始化二维切片
	// go语言中没有合适的方法，只能手动初始化
	// 二维数组的第一维用于表示横坐标，第二维表示纵坐标，建立坐标系
	// zArr := make([][]byte, numRows)
	// for i := 0; i < len(zArr); i++ {
	// 	zArr[i] = make([]byte, numRows)
	// }
	zArr := [1024][1024]byte{}
	// fmt.Println(zArr)

	sIndex := 0
	x := 0
	y := numRows - 1
	for sIndex < len(s) {
		for y = numRows - 1; y >= 0; y-- {
			if sIndex >= len(s) {
				break
			}
			// fmt.Println(sIndex)
			zArr[x][y] = s[sIndex]
			sIndex++
		}
		// 之前循环退出的时候y值需要调整
		if y < 0 {
			y++
		}
		for j := 0; j < numRows-2; j++ {
			if sIndex >= len(s) {
				break
			}
			// fmt.Println(sIndex)
			x++
			y++
			zArr[x][y] = s[sIndex]
			sIndex++
		}
		x++
		y++
	}

	// 遍历二维切片，输出图形
	// 因为以左下角作为原点，所以y轴方向是从上往下遍历的
	// 遍历y轴

		// for i := numRows - 1; i >= 0; i-- {
		// 	// 遍历x轴
		// 	for j := 0; j < numRows; j++ {
		// 		fmt.Printf("%c", zArr[j][i])
		// 	}
		// 	fmt.Println()
		// }
	res := ""
	//遍历y轴
	for i := numRows - 1; i >= 0; i-- {
		// 遍历x轴
		for j := 0; j < x; j++ {
			if zArr[j][i] == 0 {
				res += " "
			} else {
				res += string(zArr[j][i])
			}
			// fmt.Printf("%c", zArr[j][i])
		}
		// fmt.Println()
		res += "\n"
	}
	return res
}*/

// 优化:
/*
网上看到一个解法,和我的一个思路,学到了,二维切片在go里面还是直接append比较方便,初始化很麻烦
*/

func convert(s string, numRows int) string {
	var result []byte
	results := make([][]byte, numRows)
	var flag int
	for {
		for i := 0; i < numRows; i++ {
			if flag >= len(s) {
				goto Hey
			}
			results[i] = append(results[i], s[flag])
			flag++
		}
		for i := 0; i < numRows-2; i++ {
			if flag >= len(s) {
				goto Hey
			}
			results[numRows-i-2] = append(results[numRows-i-2], s[flag])
			flag++
		}
	}
Hey:
	for i := 0; i < numRows; i++ {
		result = append(result[:], results[i][:]...)
	}
	return string(result)
}
func main() {
	s := "LEETCODEISHIRING"

	s2 := "wlrbbmqbhcdarzowkkyhiddqscdxrjmowfrxsjybldbefsarcbynecdyggxxpklorellnmpapqfwkhopkmcoqhnwnkuewhsqmgbbuqcljjivswmdkqtbxixmvtrrbljptnsnfwzqfjmafadrrwsofsbcnuvqhffbsaqxwpqcacehchzvfrkmlnozjkpqpxrjxkitzyxacbhhkicqcoendtomfgdwdwfcgpxiqvkuytdlcgdewhtaciohordtqkvwcsgspqoqmsboaguwnnyqxnzlgdgwpbtrwblnsadeuguumoqcdrubetokyxhoachwdvmxxrdryxlmndqtukwagmlejuukwcibxubumenmeyatdrmydiajxloghiqfmzhl"

	// s3 := "obanbumdladpycygtrgutpdzlajzovccwcqaycfjeibclzkgsqanifmtfxsusuyqzoqxsyjwgkatllbfdesljggpdalxvjnwygvqegpmspgdcjignctxiaonazkxiyvigrbkzqwsfexiogodkjatlqioptlatjkzcllbbhthorpezfhjqkecapqsidubmecoqnsrulakerofyyrpivrkkheumyxzdzpvmhmjvpvbgkhfkyusvneiwjcijajmbzjqiwzfnuhtgoaqmukhjrpfauojwzyxyhnjfooslxorlokzlwjunaanuqzqpzqqifzoupifnwyankayhjsujukgklyckqsswtiskrzxpzackccrlxnwrxecifeuvynrrxlbqkbgkdkufpnsmaqdavhkanfxluperciinlqxoctvrindifjkaqvcgaaruryntivitnhjqcghktnvywfbocfuchoyammwwjerxoapqiwbblwbjdeygksktijuwxqsiwjhklwbtvcwgaaqfeqlqkykthgpgwkostwfhsgapkzw"
	// fmt.Println(len(s3))
	fmt.Println(convert(s, 4))
	fmt.Println(convert(s2, 6))
	// fmt.Println(convert(s3, 317))
}
