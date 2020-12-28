package main

import "fmt"

//给定一个整数 n，求以 1 ... n 为节点组成的二叉搜索树有多少种？
//
// 示例:
//
// 输入: 3
//输出: 5
//解释:
//给定 n = 3, 一共有 5 种不同结构的二叉搜索树:
//
//   1         3     3      2      1
//    \       /     /      / \      \
//     3     2     1      1   3      2
//    /     /       \                 \
//   2     1         2                 3
// Related Topics 树 动态规划
// 👍 929 👎 0

// 4:
//
//leetcode submit region begin(Prohibit modification and deletion)
func numTreesDp(n int) int {
	if n == 1 {
		return 1
	}
	dp := make([]int, n+1)
	dp[1] = 1
	for i := 2; i <= n; i++ {
		for j := 1; j <= i; j++ {
			var rightDp int
			if i-j == 0 {
				rightDp = 1
			} else {
				rightDp = dp[i-j]
			}
			var leftDp int
			if j-1 == 0 {
				leftDp = 1
			} else {
				leftDp = dp[j-1]
			}
			dp[i]+=leftDp* rightDp
		}
	}
	return dp[n]
}

func numTrees(n int) int {
	if n == 1 {
		return 1
	}
	num := 0
	dp := make([]int, n+1)
	for i := 1; i <= n; i++ {
		if i > n>>1+1 {
			num+=dp[n-i+1]
		} else {
			curNum := findRootNumTrees(i, 1, n)
			dp[i]=curNum
			num+=curNum
		}
	}
	return num
}

func findRootNumTrees(rootI int, start int, end int) int {
	if start > end {
		return 0
	}
	if start == end {
		return 1
	}
	leftNum := 0
	for i := start; i < rootI; i++ {
		curNum := findRootNumTrees(i, start, rootI-1)
		leftNum+=curNum
	}
	rightNum := 0
	for j := end; j > rootI; j-- {
		curNum := findRootNumTrees(j, rootI+1, end)
		rightNum+=curNum
	}
	if leftNum != 0 && rightNum != 0 {
		return leftNum * rightNum
	}
	return leftNum + rightNum
}
//leetcode submit region end(Prohibit modification and deletion)
func main() {
	fmt.Println(numTreesDp(19))
}
