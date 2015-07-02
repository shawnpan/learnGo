package mytest

import (
	"fmt"
)

func TestArray() {
	testArray()
	// testSlice()
}
func testArray() {
	fmt.Println("testArray-----------------------")
	var a1 [10]int = [10]int{1, 2, 3, 4}
	for i, v := range a1 {
		fmt.Printf("a1[%d]=%d,", i, v)
	}
	fmt.Println("")

	a2 := [10]int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println("a2 = ", a2, ",len=", len(a2), ",cap=", cap(a2))

	a3 := [...]int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println("a3 = ", a3, ",len=", len(a3), ",cap=", cap(a3))

	//这种是指针数组 我们看到可以直接输出指向数组的指针
	var p *[10]int = &a1
	fmt.Println("p = ", p)

	//输出类似这样[0xc080000068 0xc080000070]的地址 这就是数组指针
	x, y := 3, 5
	p1 := [...]*int{&x, &y}
	fmt.Println("p1 = ", p1)

	p2 := new([10]int)
	fmt.Println("p2 = ", p2)
}
func testSlice() {
	fmt.Println("testSlice-----------------------")
	//与数组不同的是，他不申明长度
	var s1 []int = make([]int, 3, 10) //类型，数量，容量

	fmt.Println(s1)
	fmt.Println("len = ", len(s1), ",cap = ", cap(s1)) //元素的数量和容量

	//数组转换成切片
	a := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	s2 := a[:10]
	fmt.Println("a = ", a)
	fmt.Println("s2 = ", s2)
	s3 := a[5:]
	fmt.Println("s3 = ", s3)

	//当使用append的时候，我们追加元素到切片的尾部，
	//如果我们追加的在slice容量之中的时候我们会发现，内存地址是不改变的，
	//如果我们追加的超过容量了，内存地址也就改变了
	fmt.Println("s3 len = ", len(s3), ",cap = ", cap(s3))
	fmt.Printf("s3 point = %p \n", s3)
	s3 = append(s3, 21, 22, 23)
	fmt.Println("s3 = ", s3)
	fmt.Println("s3 len = ", len(s3), ",cap = ", cap(s3))
	fmt.Printf("s3 point = %p \n", s3)

	//这是一个拷贝的函数，下边的代码是从s2拷贝到s1然后我们会看到结果是[7 8 9 4 5]
	//如果是copy(s2,s1) 我们看到的结果是[1 2 3]
	ss1 := []int{1, 2, 3, 4, 5}
	ss2 := []int{7, 8, 9}
	copy(ss1, ss2)
	fmt.Println(ss1)
	ss1 = []int{1, 2, 3, 4, 5}
	copy(ss2, ss1)
	fmt.Println(ss2)
}
