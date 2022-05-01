/*
 * @lc app=leetcode.cn id=347 lang=golang
 *
 * [347] 前 K 个高频元素
 */
package main

import (
	"container/heap"
	"fmt"
)

// @lc code=start
func topKFrequent(nums []int, k int) []int {
	seen := make(map[int]int)
	for _, v := range nums {
		seen[v]++
	}

	pq := &priorityQueue{}
	heap.Init(pq)

	for val, cnt := range seen {
		heap.Push(pq, elem{value: val, count: cnt})

		if pq.Len() > k {
			heap.Pop(pq)
		}
	}

	// myPrintf(pq)
	// return nil

	res := make([]int, 0, k)
	for pq.Len() != 0 {
		res = append(res, heap.Pop(pq).(elem).value)
	}

	return res
}

// func myPrintf(pq *priorityQueue) {
// 	fmt.Printf("len := %d\n", pq.Len())
// 	for pq.Len() != 0 {
// 		fmt.Printf("%+v\n", pq.Pop().(elem))
// 	}
// }

type elem struct {
	value, count int
}

type priorityQueue []elem

func (pq priorityQueue) Len() int           { return len(pq) }
func (pq priorityQueue) Swap(i, j int)      { pq[i], pq[j] = pq[j], pq[i] }
func (pq priorityQueue) Less(i, j int) bool { return pq[i].count < pq[j].count }

func (pq *priorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	x := old[n-1]
	*pq = old[0 : n-1]
	return x
}

func (pq *priorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(elem))
}

// @lc code=end
func main() {
	fmt.Println(topKFrequent([]int{3, 0, 1, 0}, 1))
}
