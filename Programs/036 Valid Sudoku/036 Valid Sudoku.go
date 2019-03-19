package main

import "fmt"

/*
Determine if a 9x9 Sudoku board is valid. Only the filled cells need to be validated according to the following rules:

Each row must contain the digits 1-9 without repetition.
Each column must contain the digits 1-9 without repetition.
Each of the 9 3x3 sub-boxes of the grid must contain the digits 1-9 without repetition.

A partially filled sudoku which is valid.

The Sudoku board could be partially filled, where empty cells are filled with the character '.'.

Example 1:

Input:
[
  ["5","3",'.','.',"7",'.','.','.','.'],
  ["6",'.','.',"1","9","5",'.','.','.'],
  ['.',"9","8",'.','.','.','.',"6",'.'],
  ["8",'.','.','.',"6",'.','.','.',"3"],
  ["4",'.','.',"8",'.',"3",'.','.',"1"],
  ["7",'.','.','.',"2",'.','.','.',"6"],
  ['.',"6",'.','.','.','.',"2","8",'.'],
  ['.','.','.',"4","1","9",'.','.',"5"],
  ['.','.','.','.',"8",'.','.',"7","9"]
]
Output: true
Example 2:

Input:
[
  ["8","3",'.','.',"7",'.','.','.','.'],
  ["6",'.','.',"1","9","5",'.','.','.'],
  ['.',"9","8",'.','.','.','.',"6",'.'],
  ["8",'.','.','.',"6",'.','.','.',"3"],
  ["4",'.','.',"8",'.',"3",'.','.',"1"],
  ["7",'.','.','.',"2",'.','.','.',"6"],
  ['.',"6",'.','.','.','.',"2","8",'.'],
  ['.','.','.',"4","1","9",'.','.',"5"],
  ['.','.','.','.',"8",'.','.',"7","9"]
]
Output: false
Explanation: Same as Example 1, except with the 5 in the top left corner being
    modified to 8. Since there are two 8's in the top left 3x3 sub-box, it is invalid.
Note:

A Sudoku board (partially filled) could be valid but is not necessarily solvable.
Only the filled cells need to be validated according to the mentioned rules.
The given board contain only digits 1-9 and the character '.'.
The given board size is always 9x9.


判断一个 9x9 的数独是否有效。只需要根据以下规则，验证已经填入的数字是否有效即可。

数字 1-9 在每一行只能出现一次。
数字 1-9 在每一列只能出现一次。
数字 1-9 在每一个以粗实线分隔的 3x3 宫内只能出现一次。


上图是一个部分填充的有效的数独。

数独部分空格内已填入了数字，空白格用 '.' 表示。

示例 1:

输入:
[
  ["5","3",'.','.',"7",'.','.','.','.'],
  ["6",'.','.',"1","9","5",'.','.','.'],
  ['.',"9","8",'.','.','.','.',"6",'.'],
  ["8",'.','.','.',"6",'.','.','.',"3"],
  ["4",'.','.',"8",'.',"3",'.','.',"1"],
  ["7",'.','.','.',"2",'.','.','.',"6"],
  ['.',"6",'.','.','.','.',"2","8",'.'],
  ['.','.','.',"4","1","9",'.','.',"5"],
  ['.','.','.','.',"8",'.','.',"7","9"]
]
输出: true
示例 2:

输入:
[
  ["8","3",'.','.',"7",'.','.','.','.'],
  ["6",'.','.',"1","9","5",'.','.','.'],
  ['.',"9","8",'.','.','.','.',"6",'.'],
  ["8",'.','.','.',"6",'.','.','.',"3"],
  ["4",'.','.',"8",'.',"3",'.','.',"1"],
  ["7",'.','.','.',"2",'.','.','.',"6"],
  ['.',"6",'.','.','.','.',"2","8",'.'],
  ['.','.','.',"4","1","9",'.','.',"5"],
  ['.','.','.','.',"8",'.','.',"7","9"]
]
输出: false
解释: 除了第一行的第一个数字从 5 改为 8 以外，空格内其他数字均与 示例1 相同。
     但由于位于左上角的 3x3 宫内有两个 8 存在, 因此这个数独是无效的。
说明:

一个有效的数独（部分已被填充）不一定是可解的。
只需要根据以上规则，验证已经填入的数字是否有效即可。
给定数独序列只包含数字 1-9 和字符 '.' 。
给定数独永远是 9x9 形式的。

*/

// 解法一 肯定是要遍历所有格子才能解出来
// 所以我想到的第一种解法就是遍历所有格子，而且首先遍历9个9*9小方格，这样重复遍历的部分会少一些。
// 然后遍历对应的三行三纵，每一次都用map记录数值，超过就return false
// 而且遍历时的规律规律也不好掌握？，写9次遍历，估计代码轻松上100行？
// 其实遍历规则也简单，9*9方格一层层，遍历，因为知道起点下标（左上角），也就知道终点下标（右下角）
// 然后遍历横纵的时候，方格对应的横向，纵向穿过的区域跳过

// 突然zz，做了件蠢事，数独是横方向不相同，纵方向不相同，3*3方格不相同
// 也就是说只要分别判断这三个方向相不相同就可以，我这里等于是判断3个条件同时满足的情况
/*func isValidSudoku(board [][]byte) bool {

	// 遍历9个3*3方格
	// 遍历3*3方格起始纵下标y
	for y := 0; y < 9; y += 3 {
		// 遍历3*3方格起始横下标x
		for x := 0; x < 9; x += 3 {
			// 字典存放每一次遍历的3*3方格数据
			dict := make(map[byte]int)

			for i := y; i < y+3; i++ {
				for j := x; j < x+3; j++ {
					if board[i][j] != '.' {
						dict[board[i][j]]++
						if dict[board[i][j]] > 1 {
							fmt.Println(1)
							return false
						}

					}

				}
			}
			for y1 := y; y1 < y+3; y1++ {
				for x1 := x; x1 < x+3; x1++ {
					// 复制一个9*9数组的内容
					tmpDict := make(map[byte]int)
					for index, value := range dict {
						tmpDict[index] = value
					}
					// 现在只需遍历横纵两个方向两条线
					// 遍历竖方向
					// 先遍历y2在0-y 再遍历y2在y+3 - 9 为了不便利9*9区域
					for x2, y2 := x1, 0; y2 < y; y2++ {
						if board[x2][y2] != '.' {
							tmpDict[board[x2][y2]]++
							if tmpDict[board[x2][y2]] > 1 {
								fmt.Println(2)
								return false
							}

						}
					}
					for x2, y2 := x1, y+3; y2 < 9; y2++ {
						if board[x2][y2] != '.' {
							tmpDict[board[x2][y2]]++
							if tmpDict[board[x2][y2]] > 1 {
								fmt.Println(3)
								return false
							}

						}
					}
					// 遍历横方向
					for x2, y2 := 0, y1; x2 < x; x2++ {
						if board[x2][y2] != '.' {
							tmpDict[board[x2][y2]]++
							if tmpDict[board[x2][y2]] > 1 {
								fmt.Println(4)
								return false
							}

						}
					}
					for x2, y2 := x+3, y1; x2 < 9; x2++ {
						if board[x2][y2] != '.' {
							tmpDict[board[x2][y2]]++
							if tmpDict[board[x2][y2]] > 1 {
								fmt.Println(y2, tmpDict)
								fmt.Println(5)
								return false
							}

						}
					}

				}
			}

		}
	}
	return true
}*/

// 解法一 暴力求解，遍历策略需要改一下
// 先遍历3*3方格有没有错
// 再横向遍历,纵向遍历
/*
func isValidSudoku(board [][]byte) bool {

	// 遍历9个3*3方格
	// 遍历3*3方格起始纵下标y
	for y := 0; y < 9; y += 3 {
		// 遍历3*3方格起始横下标x
		for x := 0; x < 9; x += 3 {
			// 字典存放每一次遍历的3*3方格数据
			dict := make(map[byte]int)

			for i := y; i < y+3; i++ {
				for j := x; j < x+3; j++ {
					if board[i][j] != '.' {
						dict[board[i][j]]++
						if dict[board[i][j]] > 1 {
							fmt.Println(1)
							return false
						}
					}
				}
			}
		}
	}
	for y := 0; y < 9; y++ {
		for x := 0; x < 9; x++ {

			dict := make(map[byte]int)
			// 横纵方向是分开遍历的，还是没想清楚，晚上没睡好
			// x,y相交叉的区域先判断
			// 也可以不先判断，第二次碰到的时候忽略不计即可
			// if dict[board[y][x]] != '.' {
			// 	dict[board[y][x]]++
			// }
			// 现在只需遍历横纵两个方向两条线
			// 遍历竖方向
			for x1, y1 := x, 0; y1 < 9; y1++ {
				// fmt.Println(x1, y1, board[y1][x1])
				if board[y1][x1] != '.' {
					dict[board[y1][x1]]++
					if dict[board[y1][x1]] > 1 {
						fmt.Println(2)
						return false
					}
				}
			}

			dict2 := make(map[byte]int)
			// 遍历横方向,这时过滤掉x1==x
			for x1, y1 := 0, y; x1 < 9; x1++ {
				// fmt.Println(x1, y1, dict)

				// fmt.Println(x1, y1, board[y1][x1])
				if board[y1][x1] != '.' {
					dict2[board[y1][x1]]++
					if dict2[board[y1][x1]] > 1 {
						// fmt.Println(x1, y1, dict)
						fmt.Println(3)
						return false
					}
				}

			}
		}
	}
	return true
}*/

// 解法二: 一次遍历法
// 在解法一的基础上可以优化,遍历3*3的时候,就可以把所有每行数据,每列数据,分别得出来,放到两个map里
// 这样就没有多余的遍历了
/*

func isValidSudoku(board [][]byte) bool {
	xDictSlice := make([]map[byte]int, 9)
	yDictSlice := make([]map[byte]int, 9)
	for i := range xDictSlice {
		xDictSlice[i] = make(map[byte]int)
	}
	for i := range yDictSlice {
		yDictSlice[i] = make(map[byte]int)
	}

	// 遍历9个3*3方格
	// 遍历3*3方格起始纵下标y
	for y := 0; y < 9; y += 3 {
		// 遍历3*3方格起始横下标x
		for x := 0; x < 9; x += 3 {
			// 字典存放每一次遍历的3*3方格数据
			dict := make(map[byte]int)
			for i := y; i < y+3; i++ {
				for j := x; j < x+3; j++ {
					tmpStr := board[i][j]
					if tmpStr != '.' {
						dict[tmpStr]++
						xDictSlice[j][tmpStr]++
						yDictSlice[i][tmpStr]++
						if dict[tmpStr] > 1 || xDictSlice[j][tmpStr] > 1 || yDictSlice[i][tmpStr] > 1 {
							return false
						}
					}
				}
			}
		}
	}
	return true
}
*/
// 解法二优化，其实不用mao也可以的，因为都是单个的字符，用byte数组就可以存储
// 空间和时间都能节省下来
// 这里就直接找来一位别人的代码
/*
func check3x3(mini *[][]byte, f int, p int, square byte) bool {
    count := 0
    for i:=f ; i<f+3; i++ {
        for j:=p; j<p+3; j++ {
            if square == (*mini)[i][j] && square != '.' {
                count++
            }
        }
    }
    if count >1 {
        return false
    }
    return true
}

func isValidSudoku(board [][]byte) bool {
    for i := 0; i <9; i++ {
        for j:= 0; j < 9; j++ {
            square := board[i][j]
            if (square == '.') {
                continue
            }

            // check columns
            for k:=j+1; k<9; k++ {
                if board[i][k] == square && square != '.'{
                    return false
                }
            }
            // check rows
            for m:=i+1; m<9; m++ {
                if board[m][j] == square && square != '.' {
                    return false
                }
            }
            f := (i/3) * 3
            p := (j/3) * 3

            if !(check3x3(&board, f, p, square)) {
               return false
            }

        }

    }
    return true
}*/
// 还有就算使用map，value值用bool类型就行了，判断更方便，毕竟没有计数的需要
// 时间复杂度上没有下降，但是空间使用确实减少了
// 时间复杂度O(1) 空间复杂度O(1) 因为都是固定9*9的数目，不会变化
/*
func isValidSudoku(board [][]byte) bool {
	xDictSlice := make([]map[byte]bool, 9)
	yDictSlice := make([]map[byte]bool, 9)
	for i := range xDictSlice {
		xDictSlice[i] = make(map[byte]bool)
	}
	for i := range yDictSlice {
		yDictSlice[i] = make(map[byte]bool)
	}

	// 遍历9个3*3方格
	// 遍历3*3方格起始纵下标y
	for y := 0; y < 9; y += 3 {
		// 遍历3*3方格起始横下标x
		for x := 0; x < 9; x += 3 {
			// 字典存放每一次遍历的3*3方格数据
			dict := make(map[byte]bool)
			for i := y; i < y+3; i++ {
				for j := x; j < x+3; j++ {
					tmpStr := board[i][j]
					if tmpStr != '.' {
						if dict[tmpStr] || xDictSlice[j][tmpStr] || yDictSlice[i][tmpStr] {
							return false
						}
						dict[tmpStr] = true
						xDictSlice[j][tmpStr] = true
						yDictSlice[i][tmpStr] = true
					}
				}
			}
		}
	}
	return true
}*/
/*func isValidSudoku(board [][]byte) bool {
	xDictSlice := make([][]byte, 9)
	yDictSlice := make([][]byte, 9)
	for i := range xDictSlice {
		// 0-9 +.是所有字符 ascii码最大不到60，不过我们也可以用更小的，其实10个也够了
		xDictSlice[i] = make([]byte, 60)
	}
	for i := range yDictSlice {
		yDictSlice[i] = make([]byte, 60)
	}

	// 遍历9个3*3方格
	// 遍历3*3方格起始纵下标y
	for y := 0; y < 9; y += 3 {
		// 遍历3*3方格起始横下标x
		for x := 0; x < 9; x += 3 {
			// 字典存放每一次遍历的3*3方格数据
			dict := make([]byte, 60)
			for i := y; i < y+3; i++ {
				for j := x; j < x+3; j++ {
					tmpStr := board[i][j]
					if tmpStr != '.' {
						dict[tmpStr]++
						xDictSlice[j][tmpStr]++
						yDictSlice[i][tmpStr]++
						if dict[tmpStr] > 1 || xDictSlice[j][tmpStr] > 1 || yDictSlice[i][tmpStr] > 1 {
							return false
						}
					}
				}
			}
		}
	}
	return true
}*/
func isValidSudoku(board [][]byte) bool {
	xDictSlice := make([][]byte, 9)
	yDictSlice := make([][]byte, 9)
	for i := range xDictSlice {
		// 0-9 +.是所有字符 ascii码最大不到60，不过我们也可以用更小的，其实9个也够了
		// 因为数独里面没有0，所以9也够，不过懒得改了
		xDictSlice[i] = make([]byte, 10)
	}
	for i := range yDictSlice {
		yDictSlice[i] = make([]byte, 10)
	}

	// 遍历9个3*3方格
	// 遍历3*3方格起始纵下标y
	for y := 0; y < 9; y += 3 {
		// 遍历3*3方格起始横下标x
		for x := 0; x < 9; x += 3 {
			// 字典存放每一次遍历的3*3方格数据
			dict := make([]byte, 10)
			for i := y; i < y+3; i++ {
				for j := x; j < x+3; j++ {
					tmpByte := board[i][j] - '0'
					// byte就像c语言里面一样会溢出，'.'-'0'=254
					if tmpByte != 254 {
						dict[tmpByte]++
						xDictSlice[j][tmpByte]++
						yDictSlice[i][tmpByte]++
						if dict[tmpByte] > 1 || xDictSlice[j][tmpByte] > 1 || yDictSlice[i][tmpByte] > 1 {
							return false
						}
					}
				}
			}
		}
	}
	return true
}
func main() {
	input := [][]byte{

		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}
	fmt.Println(isValidSudoku(input))
	/*var b1 byte
	var b2 byte
	b1 = '.'
	b2 = '0'
	fmt.Println(b1 - b2)*/
}
