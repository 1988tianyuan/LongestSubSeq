package main

import "fmt"

//给定一个 n × n 的二维矩阵表示一个图像。
//
// 将图像顺时针旋转 90 度。
//
// 说明：
//
// 你必须在原地旋转图像，这意味着你需要直接修改输入的二维矩阵。请不要使用另一个矩阵来旋转图像。
//
// 示例 1:
//
// 给定 matrix =
//[
//  [1,2,3],
//  [4,5,6],
//  [7,8,9]
//],
//
//原地旋转输入矩阵，使其变为:
//[
//  [7,4,1],
//  [8,5,2],
//  [9,6,3]
//]
//
//
// 示例 2:
//
// 给定 matrix =
//[
//  [ 5, 1, 9,11],
//  [ 2, 4, 8,10],
//  [13, 3, 6, 7],
//  [15,14,12,16]
//],
//
//原地旋转输入矩阵，使其变为:
//[
//  [15,13, 2, 5],
//  [14, 3, 4, 1],
//  [12, 6, 8, 9],
//  [16, 7,10,11]
//]
//
// Related Topics 数组
// 👍 725 👎 0

//[
//  [1,2,3],
//  [4,5,6],
//  [7,8,9]
//],
//[
//  [1,4,7],
//  [2,5,8],
//  [3,6,9]
//],
//leetcode submit region begin(Prohibit modification and deletion)
func rotate(matrix [][]int)  {
	for i := 0; i < len(matrix); i++ {
		for j := i; j < len(matrix); j++ {
			swapMatrix(i, j, j, i, matrix)
		}
	}
	for i := 0; i < len(matrix); i++ {
		// 逆序当前行
		reverseRow(matrix[i])
	}
}
func swapMatrix(x1 int, y1 int, x2 int, y2 int, matrix [][]int)  {
	tmp := matrix[x2][y2]
	matrix[x2][y2] = matrix[x1][y1]
	matrix[x1][y1] = tmp
}
func reverseRow(row []int) {
	start := 0
	end := len(row) - 1
	for start < end {
		tmp := row[start]
		row[start] = row[end]
		row[end] = tmp
		start++
		end--
	}
}
//leetcode submit region end(Prohibit modification and deletion)
func main() {
	matrix := [][]int{
		  {1,2,3},
		  {4,5,6},
		  {7,8,9},
	}
	rotate(matrix)
	fmt.Println(matrix)
}
