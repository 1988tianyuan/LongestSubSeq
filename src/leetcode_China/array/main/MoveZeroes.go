package main

import "fmt"

//给定一个数组 nums，编写一个函数将所有 0 移动到数组的末尾，同时保持非零元素的相对顺序。
//
// 示例:
//
// 输入: [0,1,0,3,12]
//输出: [1,3,12,0,0]
//
// 说明:
//
//
// 必须在原数组上操作，不能拷贝额外的数组。
// 尽量减少操作次数。
//
// Related Topics 数组 双指针
// 👍 895 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
func moveZeroes(nums []int)  {
	length := len(nums)
	if length <= 0 {
		return
	}
	nonZero := 0
	for i := 0; i < length; i++ {
		if nums[i] != 0 {
			swapZeroes(nonZero, i, nums)
			nonZero++
		}
	}
}

func swapZeroes(i1 int, i2 int, nums []int)  {
	tmp := nums[i2]
	nums[i2] = nums[i1]
	nums[i1] = tmp
}
//leetcode submit region end(Prohibit modification and deletion)

func main() {
	nums := []int{0,1,0,3,12}
	moveZeroes(nums)
	fmt.Println(nums)
}
