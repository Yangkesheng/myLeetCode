/*
 * @lc app=leetcode.cn id=121 lang=golang
 *
 * [121] 买卖股票的最佳时机
 */
package main

// @lc code=start
// func maxProfit(prices []int) int {
// 	now, max := 0, 0

// 	for i := 1; i < len(prices); i++ {
// 		now += prices[i] - prices[i-1]
// 		if now < 0 {
// 			now = 0
// 		} else if now > max {
// 			max = now
// 		}
// 	}

// 	return max
// }œ

// @lc code=end
