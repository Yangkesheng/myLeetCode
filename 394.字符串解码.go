/*
 * @lc app=leetcode.cn id=394 lang=golang
 *
 * [394] 字符串解码
 */
package main

import (
	"strconv"
	"strings"
)

// @lc code=start
// func decodeString(s string) string {
// 	str := ""
// 	for _, v := range s {
// 		if v == ']' {
// 			iEnd := len(str) - 1
// 			for iEnd >= 0 && !unicode.IsDigit(rune(str[iEnd])) {
// 				iEnd--
// 			}

// 			iStart := iEnd
// 			for iStart-1 >= 0 && unicode.IsDigit(rune(str[iStart-1])) {
// 				iStart--
// 			}

// 			ch := string(str[iStart : iEnd+1])
// 			d, _ := strconv.Atoi(ch)
// 			subStr := str[iEnd+2:]

// 			str = str[:iStart] + strings.Repeat(subStr, d)
// 		} else {
// 			str += string(v)
// 		}
// 	}

// 	return str
// }
func decodeString(s string) string {
	cur, _ := helper_394(s, 0)
	return cur
}

func helper_394(s string, start int) (cur string, end int) {
	k := 0
	for i := start; i < len(s); i++ {
		ch := string(s[i])

		if d, err := strconv.Atoi(string(s[i])); err == nil {
			k = k*10 + d
		} else if s[i] == '[' {
			now, end := helper_394(s, i+1)
			cur += strings.Repeat(now, k)
			k = 0
			i = end
		} else if s[i] == ']' {
			return cur, i
		} else {
			cur += ch
		}
	}

	return cur, end
}

// @lc code=end
