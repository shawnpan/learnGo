package testbase

import (
	"container/heap"
	"container/list"
	"container/ring"
	"fmt"
)

func TestContainer() {
	// testList()
	// testHeap()
	testRing()
}

func testList() {
	l := list.New()
	l.PushBack(1)
	l.PushBack(5)
	l.PushFront(3)
	l.PushFront(4)

	fmt.Println("len=", l.Len(), ",l=", l)
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}

	l.Init()
	fmt.Println("len=", l.Len(), ",l=", l)
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}

// An IntHeap is a min-heap of ints.
type IntHeap []int

func (h IntHeap) Len() int {
	return len(h)
}
func (h IntHeap) Less(i, j int) bool {
	return h[i] < h[j]
}
func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *IntHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// This example inserts several ints into an IntHeap, checks the minimum,
// and removes them in order of priority.
func testHeap() {
	h := &IntHeap{2, 1, 5}
	heap.Init(h)
	heap.Push(h, 3)
	fmt.Printf("minimum: %d\n", (*h)[0])
	for h.Len() > 0 {
		fmt.Printf("%d ", heap.Pop(h))
	}
	// Output:
	// minimum: 1
	// 1 2 3 5
}

func testRing() {
	size := 10
	r := ring.New(size)
	fmt.Println("len=", r.Len(), ",r=", r)
	num := 0
	for e := r.Prev(); e != nil; e = e.Next() {
		e.Value = num
		num++
		if num >= size {
			break
		}
	}
	num = 0
	for e := r.Prev(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
		num++
		if num >= size {
			break
		}
	}

}
