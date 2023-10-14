/*
 * @lc app=leetcode.cn id=155 lang=golang
 *
 * [155] 最小栈
 */

// @lc code=start

type Pair struct {
	Element int
	Min     int
}
type MinStack struct {
	Istack []Pair
}

func Constructor() MinStack {
	return MinStack{Istack: []Pair{}}
}

func (this *MinStack) Min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func (this *MinStack) Push(val int) {
	if len(this.Istack) == 0 {
		this.Istack = append(this.Istack, Pair{val, val})
	} else {
		this.Istack = append(this.Istack, Pair{val, this.Min(this.GetMin(), val)})
	}
}

func (this *MinStack) Pop() {
	this.Istack = this.Istack[:len(this.Istack)-1]
}

func (this *MinStack) Top() int {
	return this.Istack[len(this.Istack)-1].Element
}

func (this *MinStack) GetMin() int {
	return this.Istack[len(this.Istack)-1].Min
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
// @lc code=end

