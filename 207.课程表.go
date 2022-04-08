/*
 * @lc app=leetcode.cn id=207 lang=golang
 *
 * [207] 课程表
 */
package main

// @lc code=start
func canFinish(numCourses int, prerequisites [][]int) bool {
	graph := make(map[int][]int, numCourses)
	degree := make([]int, numCourses)

	for _, v := range prerequisites {
		graph[v[1]] = append(graph[v[1]], v[0])
		degree[v[0]] += 1
	}

	var dfs []int
	for course, d := range degree {
		if d == 0 {
			dfs = append(dfs, course)
		}
	}

	for i := 0; i < len(dfs); i++ {
		course := dfs[i]
		for _, v := range graph[course] {
			degree[v] -= 1

			if degree[v] == 0 {
				dfs = append(dfs, v)
			}
		}
	}

	return len(dfs) == numCourses
}

// @lc code=end
// func main() {
// fmt.Println(canFinish(20, [][]int{{0, 10}, {3, 18}, {5, 5}, {6, 11}, {11, 14}, {13, 1}, {15, 1}, {17, 4}}))
// fmt.Println(canFinish(2, [][]int{{1, 0}}))
// fmt.Println(canFinish(2, [][]int{{0, 10}, {0, 11}}))
// fmt.Println(canFinish(2, [][]int{{2, 1}, {1, 0}}))
// }
