package main

import (
	"container/list"
	"fmt"
	"strconv"
)

//给定一个经过编码的字符串，返回它解码后的字符串。
//
// 编码规则为: k[encoded_string]，表示其中方括号内部的 encoded_string 正好重复 k 次。注意 k 保证为正整数。
//
// 你可以认为输入字符串总是有效的；输入字符串中没有额外的空格，且输入的方括号总是符合格式要求的。
//
// 此外，你可以认为原始数据不包含数字，所有的数字只表示重复的次数 k ，例如不会出现像 3a 或 2[4] 的输入。
//
//
//
// 示例 1：
//
// 输入：s = "3[a]2[bc]"
//输出："aaabcbc"
//
//
// 示例 2：
//
// 输入：s = "3[a2[c]]"
//输出："accaccacc"
//
//
// 示例 3：
//
// 输入：s = "2[abc]3[cd]ef"
//输出："abcabccdcdcdef"
//
//
// 示例 4：
//
// 输入：s = "abc3[cd]xyz"
//输出："abccdcdcdxyz"
//
// Related Topics 栈 深度优先搜索
// 👍 624 👎 0

// 输入：s = "3[a2[c]]"
//输出："accaccacc"
//leetcode submit region begin(Prohibit modification and deletion)
const (
	leftC = 91
	rightC = 93
)
func decodeString(s string) string {
	if len(s) == 0 {
		return s
	}
	result := ""
	tmpResult := ""
	tmpNum := ""
	stack := list.New()
	for _, char := range s {
		if isSingleNum(char) != "" {
			tmpNum = tmpNum + isSingleNum(char)
			continue
		} else if tmpNum != "" {
			stack.PushFront(tmpNum)
			tmpNum = ""
		}
		if char == rightC {
			// 出栈解码
			for stack.Len() != 0 {
				e := stack.Front()
				stack.Remove(e)
				innerChar := e.Value.(string)
				tmpNum := isNum(innerChar)
				if tmpNum == -1 {
					tmpResult = innerChar + tmpResult
				} else {
					tmpTmpResult := tmpResult
					for tmpNum > 1 {
						tmpResult = tmpTmpResult + tmpResult
						tmpNum--
					}
					stack.PushFront(tmpResult)
					tmpResult = ""
					break
				}
			}
		} else {
			// 入栈
			if char != leftC {
				stack.PushFront(string(char))
			}
		}
	}
	result = tmpResult
	for stack.Len() != 0 {
		e := stack.Front()
		stack.Remove(e)
		innerChar := e.Value.(string)
		result = innerChar + result
	}
	return result
}

func isNum(char string) int {
	num,err := strconv.ParseInt(char, 10, 32)
	if err == nil {
		return int(num)
	}
	return -1
}
func isSingleNum(char int32) string {
	if char>= 48 && char <= 57 {
		return string(char)
	}
	return ""
}
//leetcode submit region end(Prohibit modification and deletion)
func main() {
	s := "3[a2[c]]"
	fmt.Println(decodeString(s))
	s2 := "abc3[cd]xyz"
	fmt.Println(decodeString(s2))
	s3 := "2[abc]3[cd]ef"
	fmt.Println(decodeString(s3))
	s4 := "100[leetcode]"
	fmt.Println(decodeString(s4))
}


