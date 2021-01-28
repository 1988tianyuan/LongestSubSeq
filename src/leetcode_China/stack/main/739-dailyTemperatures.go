package main

import (
	"container/list"
	"fmt"
)

//请根据每日 气温 列表，重新生成一个列表。对应位置的输出为：要想观测到更高的气温，至少需要等待的天数。
//如果气温在这之后都不会升高，请在该位置用 0 来代替。
//
//
// 例如，给定一个列表 temperatures = [73, 74, 75, 71, 69, 72, 76, 73]，你的输出应该是 [1, 1, 4, 2
//, 1, 1, 0, 0]。
//
// 提示：气温 列表长度的范围是 [1, 30000]。每个气温的值的均为华氏度，都是在 [30, 100] 范围内的整数。
// Related Topics 栈 哈希表
// 👍 622 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
func dailyTemperatures(T []int) []int {
	if len(T) == 0 {
		return []int{}
	}
	stask := list.New()
	tempMap := make(map[int]int)
	result := make([]int, len(T))
	result[len(T)-1] = 0
	tempMap[T[len(T)-1]] = len(T)-1
	stask.PushFront(T[len(T)-1])
	for i:=len(T)-2; i>=0; i-- {
		step := 0
		for stask.Len() != 0 {
			e := stask.Front()
			next := e.Value.(int)
			if next > T[i] {
				step = tempMap[next] - i
				break
			} else {
				stask.Remove(e)
			}
		}
		stask.PushFront(T[i])
		result[i] = step
		tempMap[T[i]] = i
	}
	return result
}
//leetcode submit region end(Prohibit modification and deletion)
func main() {
	temp := []int{73, 74, 75, 71, 69, 72, 76, 73}
	fmt.Println(dailyTemperatures(temp))

	temp2 := []int{89,62,70,58,47,47,46,76,100,70}
	fmt.Println(dailyTemperatures(temp2))
}
