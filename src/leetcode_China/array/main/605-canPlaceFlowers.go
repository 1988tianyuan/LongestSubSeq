package main

import "fmt"

//假设你有一个很长的花坛，一部分地块种植了花，另一部分却没有。可是，花卉不能种植在相邻的地块上，它们会争夺水源，两者都会死去。
//
// 给定一个花坛（表示为一个数组包含0和1，其中0表示没种植花，1表示种植了花），和一个数 n 。能否在不打破种植规则的情况下种入 n 朵花？能则返回True
//，不能则返回False。
//
// 示例 1:
//
//
//输入: flowerbed = [1,0,0,0,1], n = 1
//输出: True
//
//
// 示例 2:
//
//
//输入: flowerbed = [1,0,0,0,1], n = 2
//输出: False
//
//
// 注意:
//
//
// 数组内已种好的花不会违反种植规则。
// 输入的数组长度范围为 [1, 20000]。
// n 是非负整数，且不会超过输入数组的大小。
//
// Related Topics 贪心算法 数组
// 👍 226 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
func canPlaceFlowers(flowerbed []int, n int) bool {
	if n == 0 {
		return true
	}
	if len(flowerbed) == 1 {
		if flowerbed[0] == 1 {
			return false
		}
		return true
	}
	placed := 0
	canPlace := true
	for i := 0; i < len(flowerbed); i++ {
		if flowerbed[i] == 0 {
			if canPlace {
				if i == len(flowerbed)-1 || flowerbed[i+1] != 1 {
					flowerbed[i] = 1
					placed++
				}
				canPlace = false
			} else {
				canPlace = true
			}
		} else {
			canPlace = false
		}
		if placed == n {
			return true
		}
	}
	return false
}
func canPlaceFlowers2(flowerbed []int, n int) bool {
	if n == 0 {
		return true
	}
	if len(flowerbed) == 1 {
		if flowerbed[0] == 1 {
			return false
		}
		return true
	}
	placed := 0
	for i := 0; i < len(flowerbed); i++ {
		if flowerbed[i] == 1 {
			// 跳到i+2
			i++
		} else {
			if i == len(flowerbed)-1 || flowerbed[i+1] != 1 {
				flowerbed[i] = 1
				placed++
				i++
			}
		}
		if placed == n {
			return true
		}
	}
	return false
}
//leetcode submit region end(Prohibit modification and deletion)
func main() {
	flowerBed := []int{1,0,0,0,0,1}
	fmt.Println(canPlaceFlowers2(flowerBed, 2))
	flowerBed2 := []int{1,0,0,0,1,0,0}
	fmt.Println(canPlaceFlowers2(flowerBed2, 2))
}
