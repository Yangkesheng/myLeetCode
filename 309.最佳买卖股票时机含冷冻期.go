/*
 * @lc app=leetcode.cn id=309 lang=golang
 *
 * [309] 最佳买卖股票时机含冷冻期
 */
package main

// @lc code=start
func Max(x, y int) int {
	if x >= y {
		return x
	}
	return y
}

func maxProfit(prices []int) int {
	//持有，冷却，可以买
	hold, coolDown, canBuy := -prices[0], 0, 0

	for _, v := range prices {
		preHold, preCoolDown, preCanBuy := hold, coolDown, canBuy

		hold = Max(preHold, preCanBuy-v)
		coolDown = preHold + v
		canBuy = Max(preCoolDown, preCanBuy)
	}

	// 手上有股票一定不是最优解
	return Max(coolDown, canBuy)
}

// @lc code=end
// func main() {
// 	maxProfit([]int{1, 2, 3, 0, 2})
// }
