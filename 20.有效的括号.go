/*
 * @lc app=leetcode.cn id=20 lang=golang
 *
 * [20] 有效的括号
 */

package main

// @lc code=start
func isValid(s string) bool {
	match := map[rune]rune{
		')': '(',
		'}': '{',
		']': '[',
	}
	stack := make([]rune, 0)

	for _, v := range s {
		switch v {
		case '(', '{', '[':
			stack = append(stack, v)
		case ')', '}', ']':
			size := len(stack)
			if size > 0 && match[v] == stack[size-1] {
				stack = stack[:size-1]
			} else {
				return false
			}
		}
	}

	return len(stack) == 0
}

// @lc code=end

// func main() {
// 	isValid("()")
// }
