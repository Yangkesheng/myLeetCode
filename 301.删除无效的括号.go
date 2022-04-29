/*
 * @lc app=leetcode.cn id=301 lang=golang
 *
 * [301] 删除无效的括号
 */
package main

// @lc code=start
func isValidRes(s string) bool {
	count := 0
	for _, v := range s {
		if v == '(' {
			count++
		} else if v == ')' {
			count--

			if count < 0 {
				return false
			}
		}
	}

	return count == 0
}

func removeInvalidParentheses(s string) []string {
	ans := make([]string, 0)
	cur := map[string]struct{}{
		s: {},
	}

	for {
		for v := range cur {
			if isValidRes(v) {
				ans = append(ans, v)
			}
		}

		if len(ans) > 0 {
			return ans
		}

		next := map[string]struct{}{}
		for str := range cur {
			for k, v := range str {
				//剪枝 "((()", 只操作一次只需要删除最左一个"("，后面的情况是一样的
				if k > 0 && str[k-1] == byte(v) {
					continue
				}

				if v == '(' || v == ')' {
					next[str[:k]+str[k+1:]] = struct{}{}
				}
			}
		}

		cur = next
	}
}

// @lc code=end
