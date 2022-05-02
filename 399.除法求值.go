/*
 * @lc app=leetcode.cn id=399 lang=golang
 *
 * [399] 除法求值
 */
package main

// @lc code=start
func calcEquation(equations [][]string, values []float64, queries [][]string) []float64 {
	hash := make(map[key]float64)

	for k, v := range equations {
		hash[key{v[0], v[1]}] = values[k]
		hash[key{v[1], v[0]}] = 1 / values[k]
		hash[key{v[0], v[0]}] = 1
		hash[key{v[1], v[1]}] = 1
	}

	for k, v := range hash {
		dfs_399(k.a, k.b, v, hash)
	}

	res := make([]float64, len(queries))
	for k, v := range queries {
		if value, ok := hash[key{v[0], v[1]}]; !ok {
			res[k] = -1
		} else {
			res[k] = value
		}
	}

	return res
}

type key struct {
	a, b string
}

func dfs_399(a, b string, value float64, hash map[key]float64) {
	for k, v := range hash {
		if b == k.a && a != k.b {
			if _, ok := hash[key{a, k.b}]; !ok {
				hash[key{a, k.b}] = value * v
				dfs_399(a, k.b, value*v, hash)
			}
		}
	}
}

// @lc code=end
