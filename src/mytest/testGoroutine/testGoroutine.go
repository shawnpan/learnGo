package testGoroutine

import (
	"fmt"
)

var temp = 10
var ch chan int = make(chan int)

func TestGoroutine() {

	fmt.Println("testGoroutine")
	go func1()
	go func2()
	for i := 0; i < 2; i++ {
		<-ch
	}
}

func func1() {
	temp++
	fmt.Println(temp)
	ch <- 1
}

func func2() {
	temp--
	fmt.Println(temp)
	ch <- 1
}
