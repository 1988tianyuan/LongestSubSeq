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
	return checkExists(board, word, 0, 0, 0, mark)
}

func checkExists(board [][]byte, word string, x int, y int, position int, mark [][]int) bool {
	//time.Sleep(time.Duration(100) * time.Millisecond)
	if (y >= len(board) || y < 0) || (x >= len(board[0]) || x < 0) || position >= len(word){
		// 达到边界返回false
		return false
	}
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
	}
	mark[y][x] = 1
	d1 := checkExists(board, word, x + 1, y, position, mark)
	if d1 == true {
		return true
	}
	d2 := checkExists(board, word, x - 1, y, position, mark)
	if d2 == true {
		return true
	}
	d3 := checkExists(board, word, x, y + 1, position, mark)
	if d3 == true {
		return true
	}
	d4 := checkExists(board, word, x, y - 1, position, mark)
	if d4 == true {
		return true
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
	fmt.Println(exist(board, "ASD"))
}
