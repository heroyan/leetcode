package problems

import (
	"container/heap"
)

type ListNode struct {
	word string
	Val int
}

type ListNodeHeap []*ListNode

// https://leetcode-cn.com/problems/top-k-frequent-words/
func topKFrequent(words []string, k int) []string {
	wcntMap := map[string]int{}
	for _, str := range words {
		wcntMap[str] += 1
	}
	listHeap := &ListNodeHeap{}
	heap.Init(listHeap)
	for key, v := range wcntMap {
		if listHeap.Len() < k {
			heap.Push(listHeap, &ListNode{word: key, Val: v})
		} else if (*listHeap)[0].Val < v || ((*listHeap)[0].Val == v && (*listHeap)[0].word > key) {
			heap.Pop(listHeap)
			heap.Push(listHeap, &ListNode{word: key, Val: v})
		}
	}
	ret := []string{}
	for listHeap.Len() > 0 {
		x := heap.Pop(listHeap)
		ret = append([]string{x.(*ListNode).word}, ret...)
	}

	return ret
}


func (obj ListNodeHeap) Len() int {
	return len(obj)
}

func (obj ListNodeHeap) Less(i, j int) bool {
	if obj[i].Val == obj[j].Val {
		// 这里特别注意，越大的越容易淘汰，因为需要的结果是按字母顺序的，在出现次数相同的情况下，小字母的需要更容易保留
		return obj[i].word > obj[j].word
	}
	return obj[i].Val < obj[j].Val
}

func (obj ListNodeHeap) Swap(i, j int) {
	obj[i], obj[j] = obj[j], obj[i]
}

func (h *ListNodeHeap) Pop() interface{} {
	obj := *h
	n := len(obj)
	x := obj[n-1]
	*h = obj[0:n-1]

	return x
}

func (h *ListNodeHeap) Push(x interface{}) { // 绑定push方法，插入新元素
	*h = append(*h, x.(*ListNode))
}