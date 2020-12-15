package main

import (
	"fmt"
)

//给定一个仅包含数字 2-9 的字符串，返回所有它能表示的字母组合。
//
// 给出数字到字母的映射如下（与电话按键相同）。注意 1 不对应任何字母。
//
//
//
// 示例:
//
// 输入："23"
//输出：["ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"].
//
//
// 说明:
//尽管上面的答案是按字典序排列的，但是你可以任意选择答案输出的顺序。
// Related Topics 深度优先搜索 递归 字符串 回溯算法
// 👍 1041 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
func letterCombinations(digits string) []string {
	result := make([]string, 0)
	if len(digits) == 0 {
		return result
	}
	letterMap := make(map[uint8][]string)
	letterMap['2'] = []string{"a", "b", "c"}
	letterMap['3'] = []string{"d", "e", "f"}
	letterMap['4'] = []string{"g", "h", "i"}
	letterMap['5'] = []string{"j", "k", "l"}
	letterMap['6'] = []string{"m", "n", "o"}
	letterMap['7'] = []string{"p", "q", "r", "s"}
	letterMap['8'] = []string{"t", "u", "v"}
	letterMap['9'] = []string{"w", "x", "y", "z"}
	result = combine(digits, 0, "", result, letterMap)
	return result
}

func combine(digits string, layerI int, tmp string, result []string, letterMap map[uint8][]string) []string {
	if layerI < len(digits) - 1 {
		layer := letterMap[digits[layerI]]
		for _, c := range layer {
			result = combine(digits, layerI + 1, tmp + c, result, letterMap)
		}
	} else {
		// 最后一层
		layer := letterMap[digits[layerI]]
		for _, c := range layer {
			result = append(result, tmp+c)
		}
	}
	return result
}
//leetcode submit region end(Prohibit modification and deletion)
func main() {
	fmt.Println(letterCombinations("23"))
}
