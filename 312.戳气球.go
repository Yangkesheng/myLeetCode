/*
 * @lc app=leetcode.cn id=312 lang=golang
 *
 * [312] 戳气球
 */
package main

// @lc code=start
func maxCoins(nums []int) int {

	//create a new slice val to store the input nums
	val := make([]int, 0)
	val = append(val, 1) //nums[-1]
	val = append(val, nums...)
	val = append(val, 1) //nums[n]
	n := len(val)

	//dp[i][j] - the max coins can collect between position i and j (both exclusive)
	dp := make([][]int, len(nums)+2)
	for i := 0; i < len(nums)+2; i++ {
		dp[i] = make([]int, len(dp)+2)
	}

	for i := len(nums); i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			for k := i + 1; k < j; k++ {
				dp[i][j] = max312(dp[i][j],
					dp[i][k]+dp[k][j]+val[i]*val[k]*val[j])
			}
		}
	}

	return dp[0][n-1]
}

// func myPrint(v [][]int) {
// 	for _, v1 := range v {
// 		for _, v2 := range v1 {
// 			fmt.Printf("%4d", v2)
// 		}
// 		fmt.Println()
// 	}
// 	fmt.Println("\n=======================")
// }

func max312(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

// @lc code=end
// func main() {
// 	fmt.Println(maxCoins([]int{3, 1, 5, 8}))
// }
