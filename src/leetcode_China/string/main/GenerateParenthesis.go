package main

import (
	"fmt"
)

//数字 n 代表生成括号的对数，请你设计一个函数，用于能够生成所有可能的并且 有效的 括号组合。
//
//
//
// 示例：
//
// 输入：n = 3
//输出：[
//       "((()))",
//       "(()())",
//       "(())()",
//       "()(())",
//       "()()()"
//     ]
//
// Related Topics 字符串 回溯算法
// 👍 1476 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
const (
	left = "("
	right = ")"
)
func generateParenthesis(n int) []string {
	if n <= 0 {
		return []string{}
	}
	leftNum := 0
	rightNum := 0
	tmp := ""
	return generate(leftNum, rightNum, tmp, n, []string{})
}

/**
	括号合法性检查：
	1. 左括号数目小于n时，可以添加左括号
	2. 右括号数目小于左括号数目时，可以添加右括号
	3. 左右括号数目都为n时，则将字符串添加到结果里
 */
func generate(leftNum int, rightNum int, tmp string, n int, current []string) []string {
	if leftNum == n && rightNum == n {
		current = append(current, tmp)
	} else {
		if leftNum < n {
			current = generate(leftNum + 1, rightNum, tmp + left, n, current)
		}
		if rightNum < leftNum {
			current = generate(leftNum, rightNum + 1, tmp + right, n, current)
		}
	}
	return current
}

//leetcode submit region end(Prohibit modification and deletion)
func main() {
	fmt.Println(generateParenthesis(3))
}
