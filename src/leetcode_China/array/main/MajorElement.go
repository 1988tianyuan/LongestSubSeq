package main

import "fmt"

//给定一个大小为 n 的数组，找到其中的多数元素。多数元素是指在数组中出现次数大于 ⌊ n/2 ⌋ 的元素。
//
// 你可以假设数组是非空的，并且给定的数组总是存在多数元素。
//
//
//
// 示例 1:
//
// 输入: [3,2,3]
//输出: 3
//
// 示例 2:
//
// 输入: [2,2,1,1,1,2,2]
//输出: 2
//
// Related Topics 位运算 数组 分治算法
// 👍 822 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
func majorityElement1(nums []int) int {
	most := 0
	mostE := 0
	countMap := make(map[int]int)
	for _,e := range nums {
		count := countMap[e] + 1
		countMap[e] = count
		if count > most {
			mostE = e
			most = count
		}
	}
	return mostE
}
//leetcode submit region end(Prohibit modification and deletion)
//最优解
func majorityElement2(nums []int) int {


}

func main() {
	nums := []int{3,3,4}
	fmt.Println(majorityElement1(nums))
}