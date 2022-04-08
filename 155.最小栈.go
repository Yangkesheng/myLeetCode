/*
 * @lc app=leetcode.cn id=155 lang=golang
 *
 * [155] 最小栈
 */
package main

// @lc code=start
// type MinStack struct {
// 	stack []int
// 	min   []int
// }

// func Constructor() MinStack {
// 	return MinStack{
// 		stack: make([]int, 0),
// 		min:   make([]int, 0),
// 	}

// }

// func (this *MinStack) Push(val int) {
// 	this.stack = append(this.stack, val)

// 	l := len(this.min)
// 	if l == 0 || this.min[l-1] >= val {
// 		this.min = append(this.min, val)
// 	} else {
// 		for k, v := range this.min {
// 			if val <= v {
// 				this.min = append(this.min[:k], append([]int{val}, this.min[k:]...)...)
// 				break
// 			}
// 		}
// 	}
// }

// func (this *MinStack) Pop() {
// 	pop := this.stack[len(this.stack)-1]
// 	this.stack = this.stack[:len(this.stack)-1]

// 	for k, v := range this.min {
// 		if v == pop {
// 			this.min = append(this.min[:k], this.min[k+1:]...)
// 			break
// 		}
// 	}
// }

// func (this *MinStack) Top() int {
// 	return this.stack[len(this.stack)-1]
// }

// func (this *MinStack) GetMin() int {
// 	return this.min[len(this.min)-1]
// }

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
// @lc code=end
// func main() {
// 	min := Constructor()

// 	min.Push(512)
// 	min.Push(-1024)
// 	min.Push(-1024)
// 	min.Push(512)
// 	min.Pop()
// 	fmt.Println(min.GetMin())
// 	min.Pop()
// 	fmt.Println(min.GetMin())
// 	min.Pop()
// 	fmt.Println(min.GetMin())
// }
