package main

/*
Write a program to solve a Sudoku puzzle by filling the empty cells.

A sudoku solution must satisfy all of the following rules:

Each of the digits 1-9 must occur exactly once in each row.
Each of the digits 1-9 must occur exactly once in each column.
Each of the the digits 1-9 must occur exactly once in each of the 9 3x3 sub-boxes of the grid.
Empty cells are indicated by the character '.'.


A sudoku puzzle...


...and its solution numbers marked in red.

Note:

The given board contain only digits 1-9 and the character '.'.
You may assume that the given Sudoku puzzle will have a single unique solution.
The given board size is always 9x9.


编写一个程序，通过已填充的空格来解决数独问题。

一个数独的解法需遵循如下规则：

数字 1-9 在每一行只能出现一次。
数字 1-9 在每一列只能出现一次。
数字 1-9 在每一个以粗实线分隔的 3x3 宫内只能出现一次。
空白格用 '.' 表示。



一个数独。



答案被标成红色。

Note:

给定的数独序列只包含数字 1-9 和字符 '.' 。
你可以假设给定的数独只有唯一解。
给定数独永远是 9x9 形式的。
*/

// 解法一
// 上一题我们已经练习了判断数独是否有效的方法，那么在此之上填数独
// 遍历一次
// 我们需要3个byte数组，分别存放每一行，每一列，每一个3*3宫已有的数字，遍历一次就能全部填好
// 还需要一个数据结构，比如说一个map unsolved，存放还没填的位置，方便取出来，如果填了，那么我们就delete掉，放到一个新map里比如叫solved
// 填之前，遍历所有没有填过的点，根据每行每列，每个3*3判断哪个点可以填的可能性，map相应的count++，最后遍历map，挑一个可选数字最少的填
// 循环下去
// 如果遍历到一个无法填字的情况，那么就停止这种情况
// 如果遍历到unsloved没有数据，那么说明填好了，solved就是此时的答案，填到原来的数组上返回即可
// 综上，感觉是一个不断重复的操作，用递归实现可能会比较方便
// 循环也可以 ，退出条件就是 unsolved里的数据被delete完的时候

func solveSudoku(board [][]byte) {

}
func main() {
	/*
		var b1 byte
		var b2 byte
		b1 = '.'
		b2 = '1'
		fmt.Println(b1 - b2)//253*/
}
