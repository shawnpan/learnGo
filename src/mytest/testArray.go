package mytest

import (
	"fmt"
)

func TestArray() {
	fmt.Println("TestMap-----------------------")
	var a [10]int = [10]int{1, 2, 3, 4}
	for i, v := range a {
		fmt.Printf("a[%d]=%d,", i, v)
	}
	fmt.Println("")

	a1 := a[1:3]
	fmt.Println(a1)

	a1[1] = 99
	fmt.Println(a)
	fmt.Println(a1)

	testSlice()
}
func testSlice() {
	var s1 []int
	s1 = make([]int, 3, 10)
	fmt.Println(s1)
}
