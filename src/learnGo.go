package main

import (
	"fmt"
	_ "mytest/testGUI"
	_ "mytest/testGoroutine"
	"mytest/testIO"
	_ "mytest/testTimer"
	_ "mytest/testalgorithm"
	_ "mytest/testbase"
	_ "mytest/testdb"
)

const MAX_NUM int = 100
const MIN_NUM = 0
const MIDDLE_NUM = 50
const (
	SUNDAY = 1
	MONDAY
)

func main() {

	//	testbase.TestMap()
	//	testbase.TestArray()
	//	testbase.TestStruct()
	//	testbase.TestTime()
	//	testbase.TestString()
	//	testbase.TestStrings()
	//	testbase.TestContainer()
	//	testIO.TestFile()
	//	testIO.TestFileServer()
	//	testbase.TestOS()
	testIO.TestWeb()

	// testbase.TestIoUtil()
	// testbase.TestIni()
	// testbase.TestMysql()
	// testbase.TestJSON()
	// testbase.TestRand()
	// testbase.TestMath()
	// testbase.TestMyNet()
	// testbase.TestGet()
	// testbase.TestNewRequest()
	// testio.TestBufio()
	// testtime.TestTimer()
	// testalgorithm.TestSort()
	// testGoroutine.TestGoroutine()
	// testbase.TestChan()
	//	testGUI.TestGUI

	//	var a int
	//	fmt.Scanf("%d", &a)
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
