package main

import (
	"container/heap"
)

// 347. 前 K 个高频元素
// link: https://leetcode-cn.com/problems/top-k-frequent-elements/

type Node struct {
	val int
	times int
}
type NodeHeap []*Node
func (h NodeHeap) Len() int {
	return len(h)
}
func (h NodeHeap) Less(i, j int) bool {
	return h[i].times < h[j].times
}
func (h NodeHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h *NodeHeap) Push(x interface{}) {
	*h = append(*h, x.(*Node))
}
func (h *NodeHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}
func topKFrequent(nums []int, k int) []int {
	if k == 0 || len(nums) == 0 {
		return make([]int, 0)
	}
	// 初始化 map
	m := make(map[int]int)
	for _, v := range nums {
		m[v] = m[v]+1
	}
	h := &NodeHeap{}
	topk := min(k, len(m))
	size := 0
	for k, v := range m {
		if size < topk {
			heap.Push(h, &Node{
				val:   k,
				times: v,
			})
			size++
		} else {
			if v > (*h)[0].times {
				heap.Pop(h)
				heap.Push(h, &Node{
					val:   k,
					times: v,
				})
			}
		}
	}
	res := make([]int, 0, topk)
	for i := 0; i < topk; i++ {
		res = append(res, heap.Pop(h).(*Node).val)
	}
	return res
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}