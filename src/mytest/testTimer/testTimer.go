package testTimer

import (
	"fmt"
	"time"
)

func TestTimer() {
	testAfterFunc()
	testTimerTicker()
}

// 生成一个以每xx时间为间隔的定时器
func testTimerTicker() {
	t := time.NewTicker(time.Second * 1)
	i := 0
	for {
		select {
		case <-t.C:
			i++
			go tf2()
		}
		if i > 5 {
			break
		}
	}
}

func testAfterFunc() {
	fmt.Println("testAfterFunc ---------------------------", time.Now())
	time.AfterFunc(time.Second*2, tf1)
	fmt.Println("final -----------------------------------", time.Now())
}

func tf1() {
	fmt.Println("tf1", time.Now())
}
func tf2() {
	fmt.Println("tf2", time.Now())
}
