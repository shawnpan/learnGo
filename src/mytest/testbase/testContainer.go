package testbase

import (
	"container/heap"
	"container/list"
	"container/ring"
	"fmt"
)

func TestContainer() {
	testList()
	testRing()
	testHeap()
}

func testList() {
	fmt.Println("testList------------------------------------------------------")
	// 声明初始化
	l := list.New()
	// 在后端添加数据
	l.PushBack(1)
	l.PushBack(5)
	// 在前端添加数据
	l.PushFront(3)
	l.PushFront(4)

	fmt.Println("len=", l.Len(), ",l=", l)
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Print(e.Value, ",")
	}
	fmt.Println("")

	// 初始化、清空操作
	l.Init()
	fmt.Println("len=", l.Len(), ",l=", l)
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}

func testRing() {
	fmt.Println("testRing -----------------------------------------------------")
	// 循环列表
	size := 10
	// 声明和初始化
	r := ring.New(size)

	fmt.Println("len=", r.Len(), ",r=", r, ",r.Prev()=", r.Prev())

	for num, e := 0, r.Prev(); e != nil; e = e.Next() {
		e.Value = num
		num++
		if num >= size {
			break
		}
	}
	for num, e := 0, r.Prev(); e != nil; e = e.Next() {
		fmt.Println("num =", num, ",e.value=", e.Value)
		num++
		if num >= size+1 {
			break
		}
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
	fmt.Println("testHeap------------------------------------------------------")
	// 优先级队列
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
