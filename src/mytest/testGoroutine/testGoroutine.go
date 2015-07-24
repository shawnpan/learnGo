package testGoroutine

import (
	"fmt"
	"runtime"
)

var temp = 10
var ch chan int = make(chan int)

func TestGoroutine() {
	numCpu := runtime.NumCPU()
	fmt.Println("cpu num : ", numCpu)
	runtime.GOMAXPROCS(numCpu)
	size := 10
	for i := 0; i < size; i++ {
		go func1()
	}
	for i := 0; i < 2*size; i++ {
		go func2()
	}
	for i := 0; i < 3*size; i++ {
		<-ch
	}
	fmt.Println(temp)
	fmt.Println(1 << 20)
}

func func1() {
	temp++
	ch <- 1
}

func func2() {
	temp--
	ch <- 1
}
