/*
 * @lc app=leetcode.cn id=17 lang=golang
 *
 * [17] 电话号码的字母组合
 */
package main

// @lc code=start

var phoneNum = map[byte]string{
	'2': "abc",
	'3': "def",
	'4': "ghi",
	'5': "jkl",
	'6': "mno",
	'7': "pqrs",
	'8': "tuv",
	'9': "wxyz",
}

func letterCombinations(digits string) []string {
	res := make([]string, 0)
	if len(digits) == 0 {
		return res
	}

	cur := ""

	var dfs func(digits string)
	dfs = func(digits string) {
		str := phoneNum[digits[0]]

		for _, v := range str {
			cur += string(v)

			if len(digits) > 1 {
				dfs(digits[1:])
			} else {
				res = append(res, cur)
			}

			cur = cur[:len(cur)-1]
		}

	}

	dfs(digits)

	return res
}

// @lc code=end

func letterCombinations1(digits string) []string {
	m := map[byte][]string{
		'0': {"0"},
		'1': {"1"},
		'2': {"a", "b", "c"},
		'3': []string{"d", "e", "f"},
		'4': {"g", "h", "i"},
		'5': []string{"j", "k", "l"},
		'6': []string{"m", "n", "o"},
		'7': []string{"p", "q", "r", "s"},
		'8': []string{"t", "u", "v"},
		'9': []string{"w", "x", "y", "z"},
	}

	var res []string
	for i := 0; i < len(digits); i++ {
		digit := digits[i]

		reslen := len(res)
		if reslen == 0 {
			res = m[digit]
			continue
		}

		for j := 0; j < reslen; j++ {
			for _, ch := range m[digit] {
				res = append(res, res[j]+ch)
			}
		}
		res = res[reslen:]
	}
	return res
}

// func main() {
// 	fmt.Println(letterCombinations1(""))
// }
