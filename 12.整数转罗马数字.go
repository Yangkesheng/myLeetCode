/*
 * @lc app=leetcode.cn id=12 lang=golang
 *
 * [12] 整数转罗马数字
 */
package main

// @lc code=start
func intToRoman(num int) string {
	// cache := map[int]string{
	// 	1000: "M",
	// 	900:  "CM",
	// 	500:  "D",
	// 	400:  "CD",
	// 	100:  "C",
	// 	90:   "XC",
	// 	50:   "L",
	// 	40:   "XL",
	// 	10:   "X",
	// 	9:    "IX",
	// 	5:    "V",
	// 	4:    "IV",
	// 	1:    "I",
	// }

	n := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	value := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	res := ""
	for num > 0 {
		for k, v := range n {
			if num >= v {
				num -= v
				res += value[k]
				break
			}
		}
	}

	return res
}

// @lc code=end

// func main() {
// 	fmt.Println(intToRoman(651))
// }
