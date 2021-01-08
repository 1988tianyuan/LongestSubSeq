package main

import "fmt"

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
	heapfy(nums)
	i := 0
	for {
		nums[0], nums[len(nums)-i-1] = nums[len(nums)-i-1], nums[0]
		sinkDown(0, nums, len(nums)-i-1)
		i++
		if i == k {
			break
		}
	}
	fmt.Println(nums)
	return nums[len(nums) - k]
}

func heapfy(nums []int) {
	for i := len(nums)-1; i >= 0; i-- {
		floatUp(i, nums)
	}
}

func sinkDown(i int, nums []int, end int) {
	leftI := i*2 + 1
	rightI := i*2 + 2
	if leftI < end && nums[leftI] > nums[i] {
		nums[i], nums[leftI] = nums[leftI], nums[i]
		sinkDown(leftI, nums, end)
	}
	if rightI < end && nums[rightI] > nums[i] {
		nums[i], nums[rightI] = nums[rightI], nums[i]
		sinkDown(rightI, nums, end)
	}
}

func floatUp(i int, nums []int) {
	if i <= 0 {
		return
	}
	rootI := (i - 1)/2
	if rootI >= 0 && nums[rootI] < nums[i] {
		nums[i], nums[rootI] = nums[rootI], nums[i]
		sinkDown(i, nums, len(nums))
		floatUp(rootI, nums)
	}
}
//leetcode submit region end(Prohibit modification and deletion)
func main() {
	nums := []int{3,2,1,5,6,4}
	fmt.Println(findKthLargest(nums, len(nums)))

	nums2 := []int{3,-1,4,1,-5,9,-2,6,5}
	fmt.Println(findKthLargest(nums2, len(nums2)))

	nums3 := []int{8,14,-4,16,-3,-12,-8,-5,3,-3,-1,-8,-5,13,16,-4,17}
	fmt.Println(findKthLargest(nums3, len(nums3)))

	nums4 := []int{2,7,1,-8,-2,-8,-1,8,-2,-8,-4}
	fmt.Println(findKthLargest(nums4, len(nums4)))

	nums5 := []int{-63,-14,-79,-9,-1,-7,-54,-78,-83,8,-7,74,12,86,-78}
	fmt.Println(findKthLargest(nums5, len(nums5)))

}
