package main

import "fmt"

//用以太网线缆将 n 台计算机连接成一个网络，计算机的编号从 0 到 n-1。线缆用 connections 表示，其中 connections[i] = [
//a, b] 连接了计算机 a 和 b。
//
// 网络中的任何一台计算机都可以通过网络直接或者间接访问同一个网络中其他任意一台计算机。
//
// 给你这个计算机网络的初始布线 connections，你可以拔开任意两台直连计算机之间的线缆，并用它连接一对未直连的计算机。请你计算并返回使所有计算机都连
//通所需的最少操作次数。如果不可能，则返回 -1 。
//
//
//
// 示例 1：
//
//
//
// 输入：n = 4, connections = [[0,1],[0,2],[1,2]]
//输出：1
//解释：拔下计算机 1 和 2 之间的线缆，并将它插到计算机 1 和 3 上。
//
//
// 示例 2：
//
//
//
// 输入：n = 6, connections = [[0,1],[0,2],[0,3],[1,2],[1,3]]
//输出：2
//
//
// 示例 3：
//
// 输入：n = 6, connections = [[0,1],[0,2],[0,3],[1,2]]
//输出：-1
//解释：线缆数量不足。
//
//
// 示例 4：
//
// 输入：n = 5, connections = [[0,1],[0,2],[3,4],[2,3]]
//输出：0
//
//
//
//
// 提示：
//
//
// 1 <= n <= 10^5
// 1 <= connections.length <= min(n*(n-1)/2, 10^5)
// connections[i].length == 2
// 0 <= connections[i][0], connections[i][1] < n
// connections[i][0] != connections[i][1]
// 没有重复的连接。
// 两台计算机不会通过多条线缆连接。
//
// Related Topics 深度优先搜索 广度优先搜索 并查集
// 👍 94 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
func makeConnected(n int, connections [][]int) int {
	if len(connections) - n + 1 < 0 {
		return -1
	}
	connectedSets := make([]*ConnectedSet, 1)
	connectedSets[0] = &ConnectedSet{}
	connectMarks := make([]int, n)
	setsNum := 0
	for i := 0; i < len(connections); i++ {
		first := connections[i][0]
		second := connections[i][1]
		if connectMarks[first] != 0 && connectMarks[second] == 0 {
			connectPos := connectMarks[first]
			connectedSet := connectedSets[connectPos]
			connectedSet.set[second] = 1
			connectedSet.lineNum++
			connectMarks[second] = connectPos
		} else if connectMarks[first] == 0 && connectMarks[second] != 0 {
			connectPos := connectMarks[second]
			connectedSet := connectedSets[connectPos]
			connectedSet.set[first] = 1
			connectedSet.lineNum++
			connectMarks[first] = connectPos
		} else if connectMarks[first] == 0 && connectMarks[second] == 0 {
			newConnectedSet := &ConnectedSet{lineNum: 1, set: map[int]int{}}
			newConnectedSet.set[first] = 1
			newConnectedSet.set[second] = 1
			connectedSets = append(connectedSets, newConnectedSet)
			connectMarks[first] = len(connectedSets) - 1
			connectMarks[second] = len(connectedSets) - 1
			setsNum++
		} else {
			userMark := connectMarks[first]
			useSet := connectedSets[userMark]
			if connectMarks[first] != connectMarks[second] {
				noUseMark := connectMarks[second]
				noUseSet := connectedSets[noUseMark]
				for k := range noUseSet.set {
					connectMarks[k] = connectMarks[first]
					useSet.set[k] = 1
				}
				useSet.lineNum+=noUseSet.lineNum
				setsNum--
				connectedSets[noUseMark] = nil
			}
			useSet.lineNum++
		}
	}
	for i := 0; i < n; i++ {
		if connectMarks[i] == 0 {
			newConnectedSet := &ConnectedSet{set: map[int]int{}}
			newConnectedSet.set[i] = 1
			connectedSets = append(connectedSets, newConnectedSet)
			setsNum++
		}
	}
	return setsNum - 1
}

type ConnectedSet struct {
	lineNum int
	set map[int]int
}
//leetcode submit region end(Prohibit modification and deletion)
func main() {
	connections := [][]int{{0,1},{0,2},{0,3},{1,2},{1,3}}
	fmt.Println(makeConnected(6, connections))

	connections2 := [][]int{{0,1},{0,2},{1,2}}
	fmt.Println(makeConnected(4, connections2))

	connections3 := [][]int{{0,1},{0,2},{3,4},{2,3}}
	fmt.Println(makeConnected(5, connections3))

	connections4 := [][]int{{1,5},{1,7},{1,2},{1,4},{3,7},{4,7},{3,5},{0,6},{0,1},{0,4},{2,6},{0,3},{0,2}}
	fmt.Println(makeConnected(12, connections4))

	connections5 := [][]int{{0,1},{0,2},{0,3},{1,3}}
	fmt.Println(makeConnected(6, connections5))
}

