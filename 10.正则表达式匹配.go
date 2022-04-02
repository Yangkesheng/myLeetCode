/*
 * @lc app=leetcode.cn id=10 lang=golang
 *
 * [10] 正则表达式匹配
 */
package main

import "fmt"

// @lc code=start

func isMatch(s string, p string) bool {
	ls, lp := len(s), len(p)
	seen := make([][]bool, lp+1)
	for i := range seen {
		seen[i] = make([]bool, ls+1)
	}

	for ip := 0; ip <= lp; ip++ {
		for is := 0; is <= ls; is++ {
			if ip == 0 && is == 0 {
				seen[ip][is] = true
				continue
			} else if is == 0 {
				seen[ip][0] = p[ip-1] == '*' && seen[ip-2][0]
				continue
			} else if ip == 0 {
				continue
			}

			if p[ip-1] == '.' || p[ip-1] == s[is-1] {
				seen[ip][is] = seen[ip-1][is-1]
			} else if p[ip-1] == '*' {
				seen[ip][is] = seen[ip-1][is] || seen[ip-2][is]

				if p[ip-2] == s[is-1] || p[ip-2] == '.' {
					seen[ip][is] = seen[ip-1][is] || seen[ip-1][is-1]

					if s[is-1] == s[is-2] {
						seen[ip][is] = seen[ip][is] || seen[ip][is-1]
					}
				}
			}
		}
	}

	for _, v := range seen {
		fmt.Println(v)
	}

	return seen[lp][ls]
}

// @lc code=end

