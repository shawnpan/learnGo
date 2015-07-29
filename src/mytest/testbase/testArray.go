package testbase

import (
	"fmt"
)

func TestArray() {
	testArray()
	testSlice()
}
func testArray() {
	// 1 Go中的数组是值类型，换句话说，如果你将一个数组赋值给另外一个数组，那么，实际上就是将整个数组拷贝一份
	// 2 如果Go中的数组作为函数的参数，那么实际传递的参数是一份数组的拷贝，而不是数组的指针。这个和C要区分开。
	// 因此，在Go中如果将数组作为函数的参数传递的话，那效率就肯定没有传递指针高了。这个是不是有点陷阱的感觉？
	// 3 array的长度也是Type的一部分，这样就说明[10]int和[20]int是不一样的。
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

	// 这种是数组指针 我们看到可以直接输出指向数组的指针
	// 可以这样理解：var p *([10]int)，这个*是针对[10]int这样的数据
	var p *[10]int = &a1
	fmt.Println("p = ", p)

	// 这就是指针数组 输出类似这样[0xc080000068 0xc080000070]的地址
	// 可以这样理解：var p1 [...](*int)，这个*是表示该数组的类型为int类型的指针
	x, y := 3, 5
	var p1 = [...]*int{&x, &y}
	fmt.Println("p1 = ", p1)

	p2 := new([10]int)
	fmt.Println("p2 = ", p2)
}
func testSlice() {
	//	1 slice是可变长的，定义完一个slice变量之后，不需要为它的容量而担心，你随时可以往slice里面加数据
	//	2 slice是一个指针而不是值。指针比值可就小多了，因此，我们将slice作为函数参数传递比将array作为函数参数传递会更有性能。
	//  3 slice是一个带有point（指向数组的指针），Len（数组中实际有值的个数），Cap（数组的容量）

	fmt.Println("testSlice-----------------------")
	//与数组不同的是，他不申明长度，array有长度slice没长度
	var s1 []int = make([]int, 3, 10) //类型，数量，容量

	fmt.Println(s1)
	fmt.Println("len = ", len(s1), ",cap = ", cap(s1)) //元素的数量和容量

	//数组转换成切片
	a := [10]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
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
