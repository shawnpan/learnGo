package main

import (
	"fmt"
	"mytest"
)

const MAX_NUM int = 100
const MIN_NUM = 0
const MIDDLE_NUM = 50
const (
	SUNDAY = 1
	MONDAY
)

func main() {

	//	mytest.TestMap()
	//	mytest.TestArray()
	//	go mytest.TestStruct()
	//	mytest.TestWeb()
	// mytest.TestTime()
	// mytest.TestString()
	// mytest.TestStrings()
	// mytest.TestFile()
	// mytest.TestFileServer()
	// mytest.TestOS()
	// mytest.TestIoUtil()
	// mytest.TestIni()
	mytest.TestMysql()

}

func baseTest() {
	fmt.Println("-----------baseTest-----------------")
	i := 11.1
	fmt.Println("Hello World!")
	fmt.Println(i)
	s := "aa"
	fmt.Println(s)
	fmt.Println(MAX_NUM)
	fmt.Println(MIN_NUM)
	fmt.Println(MIDDLE_NUM)
	fmt.Println(MONDAY)
	s1 := [10]int{1, 2, 3, 4, 5}
	fmt.Println(s1)
}
