package main

import "fmt"

//给定一个非负整数 num。对于 0 ≤ i ≤ num 范围中的每个数字 i ，计算其二进制数中的 1 的数目并将它们作为数组返回。
//
// 示例 1:
//
// 输入: 2
//输出: [0,1,1]
//
// 示例 2:
//
// 输入: 5
//输出: [0,1,1,2,1,2]
//
// 进阶:
//
//
// 给出时间复杂度为O(n*sizeof(integer))的解答非常容易。但你可以在线性时间O(n)内用一趟扫描做到吗？
// 要求算法的空间复杂度为O(n)。
// 你能进一步完善解法吗？要求在C++或任何其他语言中不使用任何内置函数（如 C++ 中的 __builtin_popcount）来执行此操作。
//
// Related Topics 位运算 动态规划
// 👍 467 👎 0

// 5
// {0,1,1,2,1,2,2,3,1,2,2,3,2}
// {0,1,2,3,4,5,6,7,8,9,10,11,12}
// 1011 -> 1100 (11 -> 12)	1011-1000=0011=3  tmp=(x - (x & (x + 1))) dp[x+1]=dp[x]-(dp[tmp]-dp[tmp+1])
//
// 11 -> 100 (3 -> 4)	x=3  tmp=1
// 1101 -> 1110 (13 -> 14)
// 1001 -> 1010
//leetcode submit region begin(Prohibit modification and deletion)
func countBits(num int) []int {
	if num == 0 {
		return []int{0}
	}
	dp := make([]int, num+1)
	dp[0] = 0
	for i := 1; i <= num; i++ {
		// 求i的1的个数
		if i % 2 == 1 {
			dp[i] = dp[i-1] + 1
		} else {
			var tmp int
			noEndOne := (i-1) & i		// 消除右边连续的0
			if noEndOne == 0 {
				dp[i] = 1
			} else {
				tmp = i-1-noEndOne
				dp[i] = dp[i-1] - (dp[tmp] - dp[tmp+1])
			}
		}
	}
	return dp
}
//leetcode submit region end(Prohibit modification and deletion)
func main() {
	fmt.Println(countBits(12))
}
