package testbase

import (
	"fmt"
)

func TestMap() {
	fmt.Println("TestMap-----------------------")
	// 声明与赋值
	var m map[int]string = make(map[int]string)

	// 存数据
	m[1] = "hello"
	m[2] = "boys"

	// 遍历
	for k, v := range m {
		fmt.Printf("m[%d]=%s,", k, v)
	}
	fmt.Println("")

	imap := make(map[string]int)
	imap["a"] = 1
	imap["b"] = 2
	imap["c"] = 3
	fmt.Println(imap)

	// 初始化 + 赋值一体化
	imap2 := map[string]string{
		"a": "aa",
		"b": "bb",
	}
	fmt.Println(imap2)

	// 查找键值是否存在
	key := "b"
	v, ok := imap[key]
	if ok {
		fmt.Println("the key-value :[", key, "=", v, "]")
	} else {
		fmt.Println("not find the key : " + key)
	}

	// 删除数据
	delete(imap, "b")
	fmt.Println(imap)

	if v, ok = imap[key]; ok {
		fmt.Println("the key-value :[", key, "=", v, "]")
	} else {
		fmt.Println("not find the key : " + key)
	}

	// 嵌套map
	m1 := make(map[int]map[int]string)
	m1[1] = make(map[int]string)
	m1[1][1] = "please"
	m1[1][2] = "help"
	m1[1][3] = "me"
	fmt.Println(m1)
	m1[1][3] = "you"
	fmt.Println(m1)

	m2 := make(map[int][4]string)
	m2[1] = [4]string{"a", "b", "c", "d"}
	fmt.Println(m2)
	m2v := m2[1]
	// 修改会失败 why?
	m2v[1] = "a"
	fmt.Println(m2)

}
