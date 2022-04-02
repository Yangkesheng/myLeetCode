/*
 * @lc app=leetcode.cn id=22 lang=golang
 *
 * [22] 括号生成
 */

package main

// @lc code=start
func generateParenthesis(n int) []string {
	res := make([]string, 0)

	dfs22(&res, n, n, "")

	return res
}

func dfs22(res *[]string, l, r int, cur string) {
	if l > r { //")("
		return
	}

	if l == 0 && r == 0 {
		*res = append(*res, cur)
	}

	if l != 0 {
		dfs22(res, l-1, r, cur+"(")
	}

	if r != 0 {
		dfs22(res, l, r-1, cur+")")
	}
}

// @lc code=end

// func main() {
// 	fmt.Println(generateParenthesis(1))
// }
