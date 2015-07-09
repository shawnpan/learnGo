package testbase

import (
	"fmt"
)

func TestMap() {
	fmt.Println("TestMap-----------------------")
	var m map[int]string = make(map[int]string)
	m[1] = "hello"
	m[2] = "boys"
	for k, v := range m {
		fmt.Printf("m[%d]=%s,", k, v)
	}
	fmt.Println("")

	imap := make(map[string]int)
	imap["a"] = 1
	imap["b"] = 2
	imap["c"] = 3
	fmt.Println(imap)
	delete(imap, "b")
	fmt.Println(imap)

	m1 := make(map[int]map[int]string)
	m1[1] = make(map[int]string)
	m1[1][1] = "please"
	m1[1][2] = "help"
	m1[1][3] = "me"
	fmt.Println(m1)

}
