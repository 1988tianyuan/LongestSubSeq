package main

import (
	"container/heap"
	"fmt"
)

//有一堆石头，每块石头的重量都是正整数。
//
// 每一回合，从中选出两块 最重的 石头，然后将它们一起粉碎。假设石头的重量分别为 x 和 y，且 x <= y。那么粉碎的可能结果如下：
//
//
// 如果 x == y，那么两块石头都会被完全粉碎；
// 如果 x != y，那么重量为 x 的石头将会完全粉碎，而重量为 y 的石头新重量为 y-x。
//
//
// 最后，最多只会剩下一块石头。返回此石头的重量。如果没有石头剩下，就返回 0。
//
//
//
// 示例：
//
// 输入：[2,7,4,1,8,1]
//输出：1
//解释：
//先选出 7 和 8，得到 1，所以数组转换为 [2,4,1,1,1]，
//再选出 2 和 4，得到 2，所以数组转换为 [2,1,1,1]，
//接着是 2 和 1，得到 1，所以数组转换为 [1,1,1]，
//最后选出 1 和 1，得到 0，最终数组转换为 [1]，这就是最后剩下那块石头的重量。
//
//
//
// 提示：
//
//
// 1 <= stones.length <= 30
// 1 <= stones[i] <= 1000
//
// Related Topics 堆 贪心算法
// 👍 120 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
func lastStoneWeight(stones []int) int {
	if len(stones) == 0 {
		return 0
	}
	heapNode := &HeapNode{stones}
	heap.Init(heapNode)
	for heapNode.Len() > 1 {
		e1 := heap.Pop(heapNode).(int)
		e2 := heap.Pop(heapNode).(int)
		remain := 0
		if e1 >= e2 {
			remain = e1 - e2
		} else {
			remain = e2 - e1
		}
		if remain != 0 {
			heap.Push(heapNode, remain)
		}
	}
	if heapNode.Len() == 0 {
		return 0
	}
	return heapNode.stones[0]
}

type HeapNode struct {
	stones []int
}

func (this *HeapNode) Push(x interface{}) {
	this.stones = append(this.stones, x.(int))
}
func (this *HeapNode) Pop() interface{} {
	e := this.stones[len(this.stones)-1]
	this.stones = this.stones[0:len(this.stones)-1]
	return e
}
func (this *HeapNode) Len() int {
	return len(this.stones)
}
func (this *HeapNode) Less(i, j int) bool {
	return this.stones[i] > this.stones[j]
}
func (this *HeapNode) Swap(i, j int) {
	this.stones[i], this.stones[j] = this.stones[j], this.stones[i]
}

//leetcode submit region end(Prohibit modification and deletion)
func main() {
	stones := []int{2,7,4,1,8,1}
	fmt.Println(lastStoneWeight(stones))
}
