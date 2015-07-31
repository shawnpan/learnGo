package main

import (
	"fmt"
	"mytest/testAlgorithm"
	_ "mytest/testGUI"
	_ "mytest/testGoroutine"
	_ "mytest/testIO"
	_ "mytest/testTimer"
	_ "mytest/testbase"
	_ "mytest/testdb"
	"runtime"
)

const MAX_NUM int = 100
const MIN_NUM = 0
const MIDDLE_NUM = 50
const (
	SUNDAY = 1
	MONDAY
)

func main() {
	numCpu := runtime.NumCPU()
	runtime.GOMAXPROCS(numCpu)
	fmt.Println("runtime set cpu num : ", numCpu)
	//	testbase.TestMap()
	//	testbase.TestArray()
	//	testbase.TestStruct()
	//	testbase.TestTime()
	//	testbase.TestString()
	//	testbase.TestStrings()
	//	testbase.TestContainer()
	//	testbase.TestInterface()
	//	testbase.TestReflect()
	//	testbase.TestOS()
	//	testIO.TestFile()
	//	testIO.TestFileServer()
	//	testIO.TestWeb()
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
	//	testTimer.TestTimer()

	testAlgorithm.TestSort()
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
