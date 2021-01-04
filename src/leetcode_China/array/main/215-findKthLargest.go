package main

//在未排序的数组中找到第 k 个最大的元素。请注意，你需要找的是数组排序后的第 k 个最大的元素，而不是第 k 个不同的元素。
//
// 示例 1:
//
// 输入: [3,2,1,5,6,4] 和 k = 2
//输出: 5
//
//
// 示例 2:
//
// 输入: [3,2,3,1,2,4,5,5,6] 和 k = 4
//输出: 4
//
// 说明:
//
// 你可以假设 k 总是有效的，且 1 ≤ k ≤ 数组的长度。
// Related Topics 堆 分治算法
// 👍 841 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
func findKthLargest(nums []int, k int) int {
	if len(nums) == 0 {
		return 0
	}

}

func heapfy(nums []int) {
	for i := 0; i < len(nums); i++ {
		leftI := leftIndex(i)
		rightId := rightIndex(i)
		if leftI < len(nums) {
			if nums[leftI] > nums[i] {

			}
		}
	}
}

func leftIndex(i int) int {
	return i<<1 + 1
}
func rightIndex(i int) int {
	return i<<1 + 2
}
//leetcode submit region end(Prohibit modification and deletion)

