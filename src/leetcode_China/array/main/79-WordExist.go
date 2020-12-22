package main

import (
	"fmt"
)

//给定一个二维网格和一个单词，找出该单词是否存在于网格中。
//
// 单词必须按照字母顺序，通过相邻的单元格内的字母构成，其中“相邻”单元格是那些水平相邻或垂直相邻的单元格。同一个单元格内的字母不允许被重复使用。
//
//
//
// 示例:
//
// board =
//[
//  ['A','B','C','E'],
//  ['S','F','C','S'],
//  ['A','D','E','E']
//]
//
//给定 word = "ABCCED", 返回 true
//给定 word = "SEE", 返回 true
//给定 word = "ABCB", 返回 false
//
//
//
// 提示：
//
//
// board 和 word 中只包含大写和小写英文字母。
// 1 <= board.length <= 200
// 1 <= board[i].length <= 200
// 1 <= word.length <= 10^3
//
// Related Topics 数组 回溯算法
// 👍 717 👎 0

//{
//  {'A','B','C','E'},
//  {'S','F','C','S'},
//  {'A','D','E','E'}
//}
//leetcode submit region begin(Prohibit modification and deletion)
func exist(board [][]byte, word string) bool {
	mark := make([][]int, len(board))
	for i := 0; i < len(mark); i++ {
		mark[i] = make([]int, len(board[i]))
	}
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[0]); j++ {
			if checkExists(board, word, j, i, 0, mark) {
				return true
			}
		}
	}
	return false
}

func checkExists(board [][]byte, word string, x int, y int, position int, mark [][]int) bool {
	if mark[y][x] == 1 {
		// 已经来过了
		return false
	}
	if board[y][x] == word[position] {
		if position == len(word) - 1 {
			mark[y][x] = 1
			return true
		}
		position++
	} else {
		return false
	}
	mark[y][x] = 1
	if x + 1 < len(board[0]) {
		d1 := checkExists(board, word, x + 1, y, position, mark)
		if d1 == true {
			return true
		}
	}
	if x - 1 >= 0 {
		d2 := checkExists(board, word, x - 1, y, position, mark)
		if d2 == true {
			return true
		}
	}
	if y + 1 < len(board) {
		d3 := checkExists(board, word, x, y + 1, position, mark)
		if d3 == true {
			return true
		}
	}
	if y - 1 >= 0 {
		d4 := checkExists(board, word, x, y - 1, position, mark)
		if d4 == true {
			return true
		}
	}
	mark[y][x] = 0
	return false
}
//leetcode submit region end(Prohibit modification and deletion)
func main() {
	board := [][]byte{
		{'A','B','C','E'},
		{'S','F','C','S'},
		{'A','D','E','E'},
	}
	fmt.Println(exist(board, "BA"))
	board2 := [][]byte{
		{'a','b'},
	}
	fmt.Println(exist(board2, "ba"))
	board3 := [][]byte{
		{'F','Y','C','E','N','R','D'},
		{'K','L','N','F','I','N','U'},
		{'A','A','A','R','A','H','R'},
		{'N','D','K','L','P','N','E'},
		{'A','L','A','N','S','A','P'},
		{'O','O','G','O','T','P','N'},
		{'H','P','O','L','A','N','O'},
	}
	fmt.Println(exist(board3, "POLAND"))
	board4 := [][]byte{
		{'a','b'},
		{'c','d'},
	}
	fmt.Println(exist(board4, "acdb"))
}
