package main

import (
	"container/list"
	"fmt"
)

//给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效。
//
// 有效字符串需满足：
//
//
// 左括号必须用相同类型的右括号闭合。
// 左括号必须以正确的顺序闭合。
//
//
// 注意空字符串可被认为是有效字符串。
//
// 示例 1:
//
// 输入: "()"
//输出: true
//
//
// 示例 2:
//
// 输入: "()[]{}"
//输出: true
//
//
// 示例 3:
//
// 输入: "(]"
//输出: false
//
//
// 示例 4:
//
// 输入: "([)]"
//输出: false
//
//
// 示例 5:
//
// 输入: "{[]}"
//输出: true
// Related Topics 栈 字符串
// 👍 2035 👎 0


//leetcode submit region begin(Prohibit modification and deletion)
func isValid(s string) bool {
	if len(s) == 0 {
		return true
	}
	smallLeft := "("[0]
	smallRight := ")"[0]
	midLeft := "["[0]
	midRight := "]"[0]
	bigLeft := "{"[0]
	bigRight := "}"[0]
	stack := list.New()
	for i := 0; i < len(s); i++ {
		char := s[i]
		if char == smallLeft || char == midLeft || char == bigLeft {
			// 左括号
			stack.PushBack(char)
		} else {
			// 右括号
			if stack.Len() == 0 {
				// 栈是空的
				return false
			}
			exists := stack.Back().Value
			stack.Remove(stack.Back())
			if char == smallRight {
				if exists != smallLeft {
					return false
				}
			}
			if char == midRight {
				if exists != midLeft {
					return false
				}
			}
			if char == bigRight {
				if exists != bigLeft {
					return false
				}
			}
		}
	}
	return stack.Len() == 0
}

func main() {
	fmt.Printf("%v", isValid("()[]{}"))
	fmt.Printf("%v", isValid("([)]"))
	fmt.Printf("%v", isValid("{[]}"))
	fmt.Printf("%v", isValid("{"))
	fmt.Printf("%v", isValid("{}"))
	fmt.Printf("%v", isValid("]"))
}
//leetcode submit region end(Prohibit modification and deletion)

