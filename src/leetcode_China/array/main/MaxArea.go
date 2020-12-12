package main

import "fmt"

//给你 n 个非负整数 a1，a2，...，an，每个数代表坐标中的一个点 (i, ai) 。在坐标内画 n 条垂直线，垂直线 i 的两个端点分别为 (i,
//ai) 和 (i, 0) 。找出其中的两条线，使得它们与 x 轴共同构成的容器可以容纳最多的水。
//
// 说明：你不能倾斜容器。
//
//
//
// 示例 1：
//
//
//
//
//输入：[1,8,6,2,5,4,8,3,7]
//输出：49
//解释：图中垂直线代表输入数组 [1,8,6,2,5,4,8,3,7]。在此情况下，容器能够容纳水（表示为蓝色部分）的最大值为 49。
//
// 示例 2：
//
//
//输入：height = [1,1]
//输出：1
//
//
// 示例 3：
//
//
//输入：height = [4,3,2,1,4]
//输出：16
//
//
// 示例 4：
//
//
//输入：height = [1,2,1]
//输出：2
//
//
//
//
// 提示：
//
//
// n = height.length
// 2 <= n <= 3 * 104
// 0 <= height[i] <= 3 * 104
//
// Related Topics 数组 双指针
// 👍 2044 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
func maxArea(height []int) int {
	if len(height) == 0 {
		return 0
	}
	length := len(height)
	biggest := 0
	leftHe := height[0]
	rightHe := height[length - 1]
	left := 0
	right := length - 1
	for {
		if left >= right {
			break
		}
		area := 0
		if height[left] < leftHe {
			// 左遍历保证升序
			left++
			continue
		}
		if height[right] < rightHe {
			// 右遍历保证升序
			right--
			continue
		}
		width := right - left
		if height[left] > height[right] {
			area = height[right] * width
		} else {
			area = height[left] * width
		}
		if area > biggest {
			biggest = area
		}
		leftHe = height[left]
		rightHe = height[right]
		if leftHe > rightHe {
			right--
		} else {
			left++
		}
	}
	return biggest
}

func main()  {
	height := []int{1,8,6,2,5,4,8,3,7}
	height2 := []int{1,1}
	height3 := []int{4,3,2,1,4}
	height4 := []int{1,2,1}
	fmt.Println(maxArea(height))
	fmt.Println(maxArea(height2))
	fmt.Println(maxArea(height3))
	fmt.Println(maxArea(height4))
}
//leetcode submit region end(Prohibit modification and deletion)

