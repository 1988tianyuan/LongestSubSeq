package main

import "fmt"

//给定一个包含红色、白色和蓝色，一共 n 个元素的数组，原地对它们进行排序，使得相同颜色的元素相邻，并按照红色、白色、蓝色顺序排列。
//
// 此题中，我们使用整数 0、 1 和 2 分别表示红色、白色和蓝色。
//
//
//
// 进阶：
//
//
// 你可以不使用代码库中的排序函数来解决这道题吗？
// 你能想出一个仅使用常数空间的一趟扫描算法吗？
//
//
//
//
// 示例 1：
//
//{1,2,0,2,2,1,0,1,1}
//输入：nums = [2,0,2,1,1,0]
//输出：[0,0,1,1,2,2]
//
//
// 示例 2：
//
//
//输入：nums = [2,0,1]
//输出：[0,1,2]
//
//
// 示例 3：
//
//
//输入：nums = [0]
//输出：[0]
//
//
// 示例 4：
//
//
//输入：nums = [1]
//输出：[1]
//
//
//
//
// 提示：
//
//
// n == nums.length
// 1 <= n <= 300
// nums[i] 为 0、1 或 2
//
// Related Topics 排序 数组 双指针
// 👍 735 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
func sortColors(nums []int)  {
	p0 := 0
	p1 := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			swapSortColors(p0, i, nums)
			p0++
			if p0 > p1 {
				p1 = p0
			}
		}
		if nums[i] == 1 {
			swapSortColors(p1, i, nums)
			p1++
		}
	}
}
func swapSortColors(i1 int, i2 int, nums []int)  {
	nums[i1], nums[i2] = nums[i2], nums[i1]
}
//leetcode submit region end(Prohibit modification and deletion)
func main() {
	nums := []int{1,2,0,2,2,1,0,1,1}
	nums2 := []int{2,0,2,1,1,0}
	sortColors(nums)
	sortColors(nums2)
	fmt.Println(nums)
	fmt.Println(nums2)
}
