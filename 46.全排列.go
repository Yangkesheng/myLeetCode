/*
 * @lc app=leetcode.cn id=46 lang=golang
 *
 * [46] 全排列
 */

package main

// @lc code=start

func permuteInter(curr, subNums []int, res *[][]int) {
	if len(subNums) == 0 {
		*res = append(*res, curr)
	}

	for k, v := range subNums {
		permuteInter(
			append(curr, v),
			append(append([]int{}, subNums[:k]...), subNums[k+1:]...),
			res)
	}
}

func permute(nums []int) [][]int {
	var res [][]int

	permuteInter([]int{}, nums, &res)

	return res
}

// @lc code=end
// func main() {
// 	fmt.Println(permute([]int{1, 2, 3}))
// }
