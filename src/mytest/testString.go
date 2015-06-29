package mytest

import (
	"fmt"
)

func TestString() {
	fmt.Println("testString -------------------")
	fmt.Println("testString -------------------")
	var s string = "abcdefg"
	fmt.Println(s)

	l := len(s)
	fmt.Println(l)

	ss := s[2:5]
	fmt.Println(ss)

}
